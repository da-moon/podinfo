// Package registry acts as a registry for all request
// handlers associated with the API's V1 Spec
// the registry stores routes in the following Key-Value
// format
// <route-prefix> -> array of routes.
// ─── EXAMPLE ────────────────────────────────────────────────────────────────────
// ["/v1"] -> [{/foo},{/health}]
// ─── USAGE ──────────────────────────────────────────────────────────────────────
// call Aggregate() function when creating your http server
// and use the returned router with your stdlib's http server object.
// ────────────────────────────────────────────────────────────────────────────────
package registry
