resource "openfaas_function" "function_test" {
  name            = "test-function"
  image           = "functions/alpine:latest"
  f_process       = "env"
  labels {
    Group       = "London"
    Environment = "Test"
  }

  limits {
    memory = "20m"
    cpu    = "100m"
  }

  env_vars {
    database_name = "${postgresql_database.function_db.name}"
  }

  annotations {
    CreatedDate = "Mon Sep  3 07:15:55 BST 2018"
  }
}