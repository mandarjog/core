load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "service.go",
    ],
    importpath = "istio.io/core/pilot/test/mock",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_glog//:go_default_library",
        "@io_istio_api//proxy/v1/config:go_default_library",
        "@io_istio_pilot//model:go_default_library",
        "@io_istio_pilot//model/test:go_default_library",
        "@io_istio_pilot//proxy:go_default_library",
        "@io_istio_pilot//test/util:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    size = "small",
    srcs = ["service_test.go"],
    importpath = "istio.io/core/pilot/test/mock",
    library = ":go_default_library",
)
