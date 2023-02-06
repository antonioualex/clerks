package users

import "clerks/domain"

func CreateRoutes(us domain.UserService) map[string]domain.RouteDef {

	uh := NewUserHandler(us)

	return map[string]domain.RouteDef{
		"/populate": {
			Methods:     []string{"POST"},
			HandlerFunc: uh.Populate,
		},
		"/clerks": {
			Methods:     []string{"GET"},
			HandlerFunc: uh.Clerks,
		},
	}

}
