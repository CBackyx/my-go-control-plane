// Copyright 2020 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
package example

import (
	"time"
	"strconv"
	"os"

	"github.com/golang/protobuf/ptypes"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	auth "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/CBackyx/my-go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	v31 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
)

const (
	UserClusterName  = "user_proxy_cluster"
	RouteName    = "local_route"
	ListenerName = "listener_0"
	ListenerPort = 10000
	UserUpstreamHost = "192.168.65.2"
	UserUpstreamPort = 8081

	TempHandlerClusterName = "temphandler_cluster"

	MiddlewareClusterName = "middleware_cluster"

)

var TempHandlerUpsteamHost = os.Getenv("TEMP_HANDLER_HOST")
var TempHandlerUpsteamPort_int, err1 = strconv.Atoi(os.Getenv("TEMP_HANDLER_HOST_PORT"))
var TempHandlerUpsteamPort = uint32(TempHandlerUpsteamPort_int)
var MiddlewareUpsteamHost = os.Getenv("MIDDLEWARE_HOST")
var MiddlewareUpsteamPort_int, err2 = strconv.Atoi(os.Getenv("MIDDLEWARE_HOST_PORT"))
var MiddlewareUpsteamPort = uint32(MiddlewareUpsteamPort_int)

func makeTLSCluster(clusterName string, upstreamHost string, upstreamPort uint32) *cluster.Cluster {
	tlsc := &auth.UpstreamTlsContext{
		Sni:                  upstreamHost,                 
	}
	mt, _ := ptypes.MarshalAny(tlsc)
	return &cluster.Cluster{
		Name:                 clusterName,
		ConnectTimeout:       ptypes.DurationProto(25 * time.Second),
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_LOGICAL_DNS},
		LbPolicy:             cluster.Cluster_ROUND_ROBIN,
		LoadAssignment:       makeEndpoint(clusterName, upstreamHost, upstreamPort),
		TransportSocket:
			&core.TransportSocket{
				Name: "envoy.transport_sockets.tls",
				ConfigType: &core.TransportSocket_TypedConfig{
					TypedConfig: mt,
				},
		},
	}
}

func makeCluster(clusterName string, upstreamHost string, upstreamPort uint32) *cluster.Cluster {
	return &cluster.Cluster{
		Name:                 clusterName,
		ConnectTimeout:       ptypes.DurationProto(25 * time.Second),
		// ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_LOGICAL_DNS},
		LbPolicy:             cluster.Cluster_ROUND_ROBIN,
		LoadAssignment:       makeEndpoint(clusterName, upstreamHost, upstreamPort),
	}
}

func makeEndpoint(clusterName string, upstreamHost string, upstreamPort uint32) *endpoint.ClusterLoadAssignment {
	return &endpoint.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []*endpoint.LocalityLbEndpoints{{
			LbEndpoints: []*endpoint.LbEndpoint{{
				HostIdentifier: &endpoint.LbEndpoint_Endpoint{
					Endpoint: &endpoint.Endpoint{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Protocol: core.SocketAddress_TCP,
									Address:  upstreamHost,
									PortSpecifier: &core.SocketAddress_PortValue{
										PortValue: upstreamPort,
									},
								},
							},
						},
					},
				},
			}},
		}},
	}
}

func makeRoute(routeName string, clusterNames []string, prefixes []string, hostRewrites []string, prefixRewrites []string, regexPatterns []string, headerPatterns []string, routeTypes []int) *route.RouteConfiguration {
	var routeList []*route.Route
	regexsub := `/\1`

	for i := 0; i < len(prefixes); i++{
		if routeTypes[i] == 0 {
			routeList = append(routeList, 
				&route.Route{
					Match: &route.RouteMatch{
						PathSpecifier: &route.RouteMatch_Prefix{
							Prefix: prefixes[i],
						},
					},
					Action: &route.Route_Route{
						Route: &route.RouteAction{
							ClusterSpecifier: &route.RouteAction_Cluster{
								Cluster: clusterNames[i],
							},
							HostRewriteSpecifier: &route.RouteAction_HostRewriteLiteral{
								HostRewriteLiteral: hostRewrites[i],
							},
							PrefixRewrite: prefixRewrites[i],
						},
					},
				},
			)
		} else if routeTypes[i] == 1 {
			routeList = append(routeList, 
				&route.Route{
					Match: &route.RouteMatch{
						PathSpecifier: &route.RouteMatch_Prefix{
							Prefix: prefixes[i],
						},
						Headers: []*route.HeaderMatcher{
							&route.HeaderMatcher{
								Name: "token",
								HeaderMatchSpecifier: &route.HeaderMatcher_ExactMatch{
									ExactMatch: headerPatterns[i], 
								},
								InvertMatch: false,
							},
						},
					},
					Action: &route.Route_Route{
						Route: &route.RouteAction{
							ClusterSpecifier: &route.RouteAction_Cluster{
								Cluster: clusterNames[i],
							},
							HostRewriteSpecifier: &route.RouteAction_HostRewriteLiteral{
								HostRewriteLiteral: hostRewrites[i],
							},
							PrefixRewrite: prefixRewrites[i],
						},
					},
				},
			)
		} else if routeTypes[i] == 2 {
			routeList = append(routeList, 
				&route.Route{
					Match: &route.RouteMatch{
						PathSpecifier: &route.RouteMatch_SafeRegex{
							SafeRegex: &v31.RegexMatcher{
								EngineType: &v31.RegexMatcher_GoogleRe2{
									GoogleRe2: &v31.RegexMatcher_GoogleRE2{},
								},
								Regex: regexPatterns[i],
							},
						},
					},
					Action: &route.Route_Route{
						Route: &route.RouteAction{
							ClusterSpecifier: &route.RouteAction_Cluster{
								Cluster: clusterNames[i],
							},
							HostRewriteSpecifier: &route.RouteAction_HostRewriteLiteral{
								HostRewriteLiteral: hostRewrites[i],
							},
							RegexRewrite: &v31.RegexMatchAndSubstitute{
								Pattern: &v31.RegexMatcher{
									EngineType: &v31.RegexMatcher_GoogleRe2{
										GoogleRe2: &v31.RegexMatcher_GoogleRE2{},
									},
									Regex: regexPatterns[i],
								},
								Substitution: regexsub,
							},
						},
					},
				},
			)			
		}
	}

	return &route.RouteConfiguration{
		Name: routeName,
		VirtualHosts: []*route.VirtualHost{{
			Name:    "local_service",
			Domains: []string{"*"},
			Routes: routeList,
		}},
	}
}

func makeHTTPListener(listenerName string, route string) *listener.Listener {
	// HTTP filter configuration
	manager := &hcm.HttpConnectionManager{
		UpgradeConfigs: []*hcm.HttpConnectionManager_UpgradeConfig{{
			UpgradeType: "websocket",
			Enabled:     &wrappers.BoolValue{Value: true},
		}},
		CodecType:  hcm.HttpConnectionManager_AUTO,
		StatPrefix: "http",
		RouteSpecifier: &hcm.HttpConnectionManager_Rds{
			Rds: &hcm.Rds{
				ConfigSource:    makeConfigSource(),
				RouteConfigName: route,
			},
		},
		HttpFilters: []*hcm.HttpFilter{{
			Name: wellknown.Router,
		}},
	}
	pbst, err := ptypes.MarshalAny(manager)
	if err != nil {
		panic(err)
	}

	return &listener.Listener{
		Name: listenerName,
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.SocketAddress_TCP,
					Address:  "0.0.0.0",
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: ListenerPort,
					},
				},
			},
		},
		FilterChains: []*listener.FilterChain{{
			Filters: []*listener.Filter{{
				Name: wellknown.HTTPConnectionManager,
				ConfigType: &listener.Filter_TypedConfig{
					TypedConfig: pbst,
				},
			}},
		}},
	}
}

func makeConfigSource() *core.ConfigSource {
	source := &core.ConfigSource{}
	source.ResourceApiVersion = resource.DefaultAPIVersion
	source.ConfigSourceSpecifier = &core.ConfigSource_ApiConfigSource{
		ApiConfigSource: &core.ApiConfigSource{
			TransportApiVersion:       resource.DefaultAPIVersion,
			ApiType:                   core.ApiConfigSource_GRPC,
			SetNodeOnFirstMessageOnly: true,
			GrpcServices: []*core.GrpcService{{
				TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
					EnvoyGrpc: &core.GrpcService_EnvoyGrpc{ClusterName: "xds_cluster"},
				},
			}},
		},
	}
	return source
}

func GenerateSnapshot(tokenMap map[string]SingleRouteInfo) cache.Snapshot {
	// Get original SnapshotFeed, some fixed pattern
	originalSnapshotFeed :=	GenerateOriginalSnapshotFeed()
	for _, sri := range tokenMap {
		originalSnapshotFeed.clusters = append(originalSnapshotFeed.routes, makeCluster(sri.clusterName, sri.routeHostRewrite, sri.hostPort))		
	}

	// Append original route set
	// routeClusterNames := []string{UserClusterName}
	// routePrefixes := []string{"/user"}
	// routeHostRewrites := []string{UserUpstreamHost}
	// routePrefixRewrites := []string{"/user"}
	// routeTypes := []int{0}
	// routeRegexes := []string{""}
	// routeHeaders := []string{""}
	routeClusterNames := []string{UserClusterName, TempHandlerClusterName, MiddlewareClusterName}
	routePrefixes := []string{"/user", "/temphandler/", "/middleware/"}
	routeHostRewrites := []string{UserUpstreamHost, TempHandlerUpsteamHost, MiddlewareUpsteamHost}
	routePrefixRewrites := []string{"/user", "/temphandler/", "/middleware/"}
	routeTypes := []int{0, 0, 0}
	routeRegexes := []string{"", "", ""}
	routeHeaders := []string{"", "", ""}

	// Append token->container route set, including header match
	for _, sri := range tokenMap {
		routeClusterNames = append(routeClusterNames, sri.clusterName)
		routePrefixes = append(routePrefixes, sri.routePrefix)
		routeHostRewrites = append(routeHostRewrites, sri.routeHostRewrite)
		routePrefixRewrites = append(routePrefixRewrites, sri.routePrefixRewrite)
		routeTypes = append(routeTypes, 1)
		routeRegexes = append(routeRegexes, "")
		routeHeaders = append(routeHeaders, sri.matchToken)	
	}

	// Append regex match, to handle unmatched token->container above
	routeClusterNames = append(routeClusterNames, TempHandlerClusterName)
	routePrefixes = append(routePrefixes, "/webide/")
	routeHostRewrites = append(routeHostRewrites, TempHandlerUpsteamHost)
	routePrefixRewrites = append(routePrefixRewrites, "")
	routeTypes = append(routeTypes, 2)
	routeRegexes = append(routeRegexes, `^/webide/[abc]lo/(.*)$`)
	routeHeaders = append(routeHeaders, "")		

	return cache.NewSnapshot(
		"1",
		[]types.Resource{}, // endpoints
		originalSnapshotFeed.clusters,
		[]types.Resource{makeRoute(RouteName, routeClusterNames, routePrefixes, routeHostRewrites, routePrefixRewrites, routeRegexes, routeHeaders, routeTypes)},
		[]types.Resource{makeHTTPListener(ListenerName, RouteName)},
		[]types.Resource{}, // runtimes
		[]types.Resource{}, // secrets
	)
}

func GenerateOriginalSnapshot() cache.Snapshot {

	// routeClusterNames := []string{UserClusterName}
	// routePrefixes := []string{"/webide/hwudhdnnsjwkzjnsnwjw/"}
	// routeHostRewrites := []string{UserUpstreamHost}
	// routePrefixRewrites := []string{"/"}
	// routeTypes := []int{0}
	// routeRegexes := []string{}
	// routeHeaders := []string{}

	// return cache.NewSnapshot(
	// 	"1",
	// 	[]types.Resource{}, // endpoints
	// 	[]types.Resource{makeCluster(UserClusterName, UserUpstreamHost, UserUpstreamPort)},
	// 	[]types.Resource{makeRoute(RouteName, routeClusterNames, routePrefixes, routeHostRewrites, routePrefixRewrites, routeRegexes, routeHeaders, routeTypes)},
	// 	[]types.Resource{makeHTTPListener(ListenerName, RouteName)},
	// 	[]types.Resource{}, // runtimes
	// 	[]types.Resource{}, // secrets
	// )

	routeClusterNames := []string{UserClusterName, TempHandlerClusterName, MiddlewareClusterName}
	routePrefixes := []string{"/user", "/temphandler/", "/middleware/"}
	routeHostRewrites := []string{UserUpstreamHost, TempHandlerUpsteamHost, MiddlewareUpsteamHost}
	routePrefixRewrites := []string{"/user", "/temphandler/", "/middleware/"}
	routeTypes := []int{0, 0, 0}
	routeRegexes := []string{"", "", ""}
	routeHeaders := []string{"", "", ""}

	return cache.NewSnapshot(
		"1",
		[]types.Resource{}, // endpoints
		[]types.Resource{
						makeCluster(UserClusterName, UserUpstreamHost, UserUpstreamPort), 
						makeCluster(TempHandlerClusterName, TempHandlerUpsteamHost, TempHandlerUpsteamPort),
						makeCluster(MiddlewareClusterName, MiddlewareUpsteamHost, MiddlewareUpsteamPort),
					},
		[]types.Resource{makeRoute(RouteName, routeClusterNames, routePrefixes, routeHostRewrites, routePrefixRewrites, routeRegexes, routeHeaders, routeTypes)},
		[]types.Resource{makeHTTPListener(ListenerName, RouteName)},
		[]types.Resource{}, // runtimes
		[]types.Resource{}, // secrets
	)
}

type SingleRouteInfo struct {
	clusterName        string
	routePrefix        string
	hostPort           uint32
	routeHostRewrite   string
	routePrefixRewrite string
	routeRegex         string
	matchToken         string
}

type SnapshotFeed struct {
	endpoints []types.Resource
	clusters  []types.Resource
	routes    []types.Resource
	listeners []types.Resource
	runtimes  []types.Resource
	secrets   []types.Resource
}

func GenerateOriginalSnapshotFeed() SnapshotFeed {
	routeClusterNames := []string{UserClusterName, TempHandlerClusterName, MiddlewareClusterName}
	routePrefixes := []string{"/user", "/temphandler/", "/middleware/"}
	routeHostRewrites := []string{UserUpstreamHost, TempHandlerUpsteamHost, MiddlewareUpsteamHost}
	routePrefixRewrites := []string{"/user", "/temphandler/", "/middleware/"}
	routeTypes := []int{0, 0, 0}
	routeRegexes := []string{"", "", ""}
	routeHeaders := []string{"", "", ""}

	return SnapshotFeed{
		[]types.Resource{}, // endpoints
		[]types.Resource{
						makeCluster(UserClusterName, UserUpstreamHost, UserUpstreamPort), 
						makeCluster(TempHandlerClusterName, TempHandlerUpsteamHost, TempHandlerUpsteamPort),
						makeCluster(MiddlewareClusterName, MiddlewareUpsteamHost, MiddlewareUpsteamPort),
					},
		[]types.Resource{makeRoute(RouteName, routeClusterNames, routePrefixes, routeHostRewrites, routePrefixRewrites, routeRegexes, routeHeaders, routeTypes)},
		[]types.Resource{makeHTTPListener(ListenerName, RouteName)},
		[]types.Resource{}, // runtimes
		[]types.Resource{}, // secrets
	}
}

func GenerateTestSnapshot(num int) cache.Snapshot {
	routeClusterNames := []string{UserClusterName, "regex_cluster"}
	routePrefixes := []string{"/webide/" + strconv.Itoa(num) + "/", "/webide/"}
	routeHostRewrites := []string{UserUpstreamHost, UserUpstreamHost}
	routePrefixRewrites := []string{"/", "/"}
	routeTypes := []int{1, 2}
	routeRegexes := []string{"", `^/webide/[abc]lo/(.*)$`}
	routeHeaders := []string{"ahjsmnvabhfvaeub", ""}


	return cache.NewSnapshot(
		"1",
		[]types.Resource{}, // endpoints
		[]types.Resource{makeCluster(UserClusterName, UserUpstreamHost, UserUpstreamPort), makeCluster("regex_cluster", UserUpstreamHost, UserUpstreamPort)},
		[]types.Resource{makeRoute(RouteName, routeClusterNames, routePrefixes, routeHostRewrites, routePrefixRewrites, routeRegexes, routeHeaders, routeTypes)},
		[]types.Resource{makeHTTPListener(ListenerName, RouteName)},
		[]types.Resource{}, // runtimes
		[]types.Resource{}, // secrets
	)
}


