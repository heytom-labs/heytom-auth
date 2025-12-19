package server

import (
	authv1 "heytom-auth/api/auth/v1"
	v1 "heytom-auth/api/helloworld/v1"
	"heytom-auth/internal/conf"
	"heytom-auth/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, user *service.UserService, role *service.RoleService, policy *service.PolicyService, application *service.ApplicationService, organization *service.OrganizationService, auth *service.AuthService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	authv1.RegisterUserHTTPServer(srv, user)
	authv1.RegisterRoleHTTPServer(srv, role)
	authv1.RegisterPolicyHTTPServer(srv, policy)
	authv1.RegisterApplicationHTTPServer(srv, application)
	authv1.RegisterOrganizationHTTPServer(srv, organization)
	authv1.RegisterAuthHTTPServer(srv, auth)
	return srv
}
