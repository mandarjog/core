load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "controller.go",
        "conversion.go",
        "serviceaccounts.go",
        "servicediscovery.go",
    ],
    importpath = "istio.io/core/pilot/platform/eureka",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_glog//:go_default_library",
        "@io_istio_pilot//model:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    size = "small",
    srcs = [
        "client_test.go",
        "controller_test.go",
        "conversion_test.go",
        "servicediscovery_test.go",
    ],
    data = glob(["testdata/*"]),
    importpath = "istio.io/core/pilot/platform/eureka",
    library = ":go_default_library",
    deps = [
        "@io_istio_pilot//model:go_default_library",
        "@io_istio_pilot//test/util:go_default_library",
    ],
)
