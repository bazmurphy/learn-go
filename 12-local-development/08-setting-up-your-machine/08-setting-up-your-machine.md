# Setting up your machine

Your machine will contain many version control _repositories_ (managed by Git, for example).

Each repository contains one or more _packages_, but will typically be a single _module_.

Each package consists of one or more _Go source files_ in a single directory.

The path to a package's directory determines its _import path_ and where it can be downloaded from if you decide to host it on a remote version control system like Github or Gitlab.

## A note on GOPATH

The $GOPATH environment variable will be set by default somewhere on your machine (typically in the home directory, `~/go`). Since we will be working in the new "Go modules" setup, you _don't need to worry about that_. If you read something online about setting up your GOPATH, that documentation is probably out of date.

These days you should _avoid_ working in the `$GOPATH/src` directory. Again, that's the old way of doing things and can cause unexpected issues, so better to just avoid it.

## Get into your workspace

Navigate to a location on your machine where you want to store some code. For example, I store all my code in `~/workspace`, then organize it into subfolders based on the remote location. For example,

`~/workspace/github.com/wagslane/go-password-validator` = [https://github.com/wagslane/go-password-validator](https://github.com/wagslane/go-password-validator)

That said, you can put your code wherever you want.
