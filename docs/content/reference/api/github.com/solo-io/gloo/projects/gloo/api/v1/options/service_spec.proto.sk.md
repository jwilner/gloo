
---
title: "service_spec.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `options.gloo.solo.io` 
#### Types:


- [ServiceSpec](#servicespec)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/v1/options/service_spec.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/options/service_spec.proto)





---
### ServiceSpec

 
Describes APIs and application-level information for services
Gloo routes to. ServiceSpec is contained within the UpstreamSpec for certain types
of upstreams, including Kubernetes, Consul, and Static.
ServiceSpec configuration is opaque to Gloo and handled by Service Options.

```yaml
"rest": .rest.options.gloo.solo.io.ServiceSpec
"grpc": .grpc.options.gloo.solo.io.ServiceSpec

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `rest` | [.rest.options.gloo.solo.io.ServiceSpec](../rest/rest.proto.sk/#servicespec) |  Only one of `rest` or `grpc` can be set. |  |
| `grpc` | [.grpc.options.gloo.solo.io.ServiceSpec](../grpc/grpc.proto.sk/#servicespec) |  Only one of `grpc` or `rest` can be set. |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->