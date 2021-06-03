provider "epcc" {
  // Can set via `EPCC_CLIENT_ID` environment variable.
  client_id = "some_client_id"

  // Can set via `EPCC_CLIENT_SECRET` environment variable.
  client_secret = "some_client_secret"

  // Can set via `EPCC_API_BASE_URL` environment variable
  api_base_url = "https://api.moltin.com/"

  // Can set via `EPCC_BETA_API_FEATURES` environment variable.
  beta_features = "account-management"
}