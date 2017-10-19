load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "controller.go",
        "monitor.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_glog//:go_default_library",
        "@io_istio_core//pilot/model:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    size = "small",
    srcs = [
        "config_test.go",
        "controller_test.go",
        "monitor_test.go",
    ],
    deps = [
        "@io_istio_core//pilot/adapter/config/memory:go_default_library",
        "@io_istio_core//pilot/model:go_default_library",
        "@io_istio_core//pilot/test/mock:go_default_library",
    ],
)
