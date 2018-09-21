provider "openfaas" {
  version = "~> 0.1"
  uri       = "http://localhost:8080"
  user_name = "admin"
  password  = "${var.provider_password}"
}

resource "openfaas_function" "function_test" {
  name            = "test-function"
  image           = "functions/alpine:latest"
  f_process       = "env"
  labels {
    Group       = "London"
    Environment = "Test"
  }

  annotations {
    CreatedDate = "Mon Sep  3 07:15:55 BST 2018"
  }
}
