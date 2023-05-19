package task

// config stock la configuration de l'application.
// C'est le bon endroit pour "pluger" des comportements externes au
// métier comme le dépot de données.
// Pas besoin de l'exposer à l'extérieur, donc un c minuscule !
var config = struct {
	repo Repository
}{}

// WithRepo permet de pluger le dépot de données
func WithRepo(r Repository) {
	config.repo = r
}
