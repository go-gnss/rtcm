## RTCM parser library implementation for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/go-gnss/rtcm)](https://goreportcard.com/report/github.com/go-gnss/rtcm)
[![Coverage Status](https://coveralls.io/repos/github/go-gnss/rtcm/badge.svg?branch=master)](https://coveralls.io/github/go-gnss/rtcm?branch=master)

Example usage in `cmd/ntriplatency`

```
ntriplatency -caster http://caster:port/mount -username username -password password

```

The biggest issue with this library is that it's not useful for constructing RTCM messages (only
really for parsing them). It can be done, but realistically an abstraction would be required so
that the developer experience isn't so painful and error prone.

An example of this is that masking is used to identify list lengths and satellite IDs in the RTCM
binary structure. Meaning if you add a satellite to an MessageMsm7.SatelliteData, you also have to
make a change to the satellite and signal masks. Ideally this would be abstracted such that the
masks are computed during Serialization based on the data in the object.

A normalized Observation object may be useful for this. This would lead to scenarios like the
following being easily implementable: input io provides a binary MSM7 1077 message which is
deserialized into a normalized Observation struct, problematic satellites are removed, then the
Observation is Serialized to output io as a binary MSM5 1075 message.
