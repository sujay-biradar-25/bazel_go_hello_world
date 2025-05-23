def go_library(name, srcs, deps=[]):
    all_srcs = srcs + [dep + ".a" for dep in deps]
    native.genrule(
        name = name,
        srcs = all_srcs,
        outs = [name + ".a"],
        cmd = "go tool compile -I . -o $(OUTS) $(SRCS)",
    )

def go_binary(name, srcs, deps=[]):
    all_deps = deps + [dep + ".a" for dep in deps]
    native.genrule(
        name = name,
        srcs = srcs + all_deps,
        outs = [name],
        cmd = "go tool link -o $(OUTS) $(SRCS) $(DEPS)",
    )