load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "istio.io/core/pilot/cmd/pilot-agent",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_golang_glog//:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library",
        "@com_github_golang_protobuf//ptypes/duration:go_default_library",
        "@com_github_spf13_cobra//:go_default_library",
        "@io_istio_api//proxy/v1/config:go_default_library",
        "@io_istio_pilot//cmd:go_default_library",
        "@io_istio_pilot//model:go_default_library",
        "@io_istio_pilot//platform:go_default_library",
        "@io_istio_pilot//proxy:go_default_library",
        "@io_istio_pilot//proxy/envoy:go_default_library",
        "@io_istio_pilot//tools/version:go_default_library",
    ],
)

go_binary(
    name = "pilot-agent",
    importpath = "istio.io/core/pilot/cmd/pilot-agent",
    library = ":go_default_library",
    linkstamp = "istio.io/pilot/tools/version",
    visibility = ["//visibility:public"],
)
