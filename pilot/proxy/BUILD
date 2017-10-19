load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "agent.go",
        "context.go",
        "net.go",
        "resolve.go",
    ],
    importpath = "istio.io/core/pilot/proxy",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_glog//:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library",
        "@com_github_hashicorp_go_multierror//:go_default_library",
        "@io_istio_api//proxy/v1/config:go_default_library",
        "@io_istio_pilot//model:go_default_library",
        "@org_golang_x_time//rate:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    size = "small",
    srcs = ["agent_test.go"],
    importpath = "istio.io/core/pilot/proxy",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    size = "small",
    srcs = ["context_test.go"],
    importpath = "istio.io/core/pilot/proxy_test",
    deps = [
        "@io_istio_pilot//model:go_default_library",
        "@io_istio_pilot//proxy:go_default_library",
        "@io_istio_pilot//test/mock:go_default_library",
    ],
)
