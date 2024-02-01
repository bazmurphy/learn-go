module someremote/bazmurphy/hellogo

go 1.21.6

// because go packages are heavily dependent on remotely hosted packages
// we need to override the location of the mystrings package to point to the local one
replace someremote/bazmurphy/mystrings v0.0.0 => ../mystrings

require someremote/bazmurphy/mystrings v0.0.0