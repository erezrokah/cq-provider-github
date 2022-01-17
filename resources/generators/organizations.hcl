service = "github"
output_directory = "../resources"
add_generate = true

description_source "graphql" {
    path = "generators/schema.docs.graphql"
}


resource "github" "" "organizations" {
    path = "github.com/google/go-github/v41/github.Organization"
    description_path_parts = ["Organization"]
    options {
        primary_keys = [
            "id"]
    }
    column "id" {
        description = "test"
    }

}