
locals {
  emissary_config_yaml = <<-EOT
    hosts:
    - name: my-host-dev
      spec:
        ambassador_id: 
        - ${local.ambassador_id}
        hostname: '*.wadexu.cloud'
        requestPolicy:
          insecure:
            action: Redirect
        tlsContext:
          name: my-tls-context
        tlsSecret:
          name: tls-secret

    mappings:
    - name: my-nginx-mapping
      spec:
        ambassador_id:
        - ${local.ambassador_id}
        hostname: dev.wadexu.cloud
        prefix: /
        service: my-nginx.nginx:80

    tlscontexts:
    - name: my-tls-context
      spec:
        ambassador_id: 
        - ${local.ambassador_id}
        hosts:
        - "*.wadexu.cloud"
        min_tls_version: v1.2
  EOT
}
