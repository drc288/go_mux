resource "mongodbatlas_project" main {
    name = "go_mux_udemy"
    org_id = var.organization_id
}


resource "mongodbatlas_cluster" "main" {
  project_id              = mongodbatlas_project.main.id
  name                    = "goMux"

  # Provider Settings "block"
  provider_name = "TENANT"
  backing_provider_name = "AZURE"
  provider_region_name = "us-central1"
  provider_instance_size_name = "M0"
}