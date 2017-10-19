load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "istio.io/core/auth/cmd/generate_csr",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_golang_glog//:go_default_library",
        "@io_istio_auth//pkg/pki/ca:go_default_library",
    ],
)

go_binary(
    name = "generate_csr",
    importpath = "istio.io/core/auth/cmd/generate_csr",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
