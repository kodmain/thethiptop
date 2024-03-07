global:
  scrape_interval:     5s
  evaluation_interval: 5s

scrape_configs:
  - job_name: 'nomad'
    metrics_path: '/v1/metrics'
    scheme: https
    params:
        format: ['prometheus']
    static_configs:
      - targets: ["{{ env "NOMAD_HOST_ADDR_nomad" }}"]
  - job_name: 'node-exporter'
    static_configs:
      - targets: ["{{ env "NOMAD_HOST_ADDR_node-exporter" }}"]
