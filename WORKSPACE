workspace(name = "bazel_go_hello_world")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Go rules
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "b2038e2de2cace18f032249cb4bb0048abf583a36369fa98f687af1b3f880b26",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.48.1/rules_go-v0.48.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.48.1/rules_go-v0.48.1.zip",
    ],
)

# Gazelle
http_archive(
    name = "bazel_gazelle",
    sha256 = "32938bda16e6700063035479063d9d24c60eda8d79fd4739563f50d331cb3209",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
    ],
)

# Load Go rules dependencies
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_rules_dependencies()

go_register_toolchains(version = "1.21.6")

gazelle_dependencies()

# External Go dependencies
go_repository(
    name = "com_github_bwmarrin_snowflake",
    importpath = "github.com/bwmarrin/snowflake",
    sum = "h1:xm67bEhkKh6ij1790JB83OujPR5CzNe8QuQqAgISZN0=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=",
    version = "v1.6.0",
) 