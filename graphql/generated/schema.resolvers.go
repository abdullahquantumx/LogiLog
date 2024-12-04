package generated

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.57

import (
	"context"
	"fmt"

	main "github.com/Shridhar2104/logilo/graphql"
)

// Email is the resolver for the email field.
func (r *accountResolver) Email(ctx context.Context, obj *main.Account) (string, error) {
	panic(fmt.Errorf("not implemented: Email - email"))
}

// PhoneNumber is the resolver for the phoneNumber field.
func (r *accountResolver) PhoneNumber(ctx context.Context, obj *main.Account) (*string, error) {
	panic(fmt.Errorf("not implemented: PhoneNumber - phoneNumber"))
}

// DefaultShippingAddress is the resolver for the defaultShippingAddress field.
func (r *accountResolver) DefaultShippingAddress(ctx context.Context, obj *main.Account) (*main.Address, error) {
	panic(fmt.Errorf("not implemented: DefaultShippingAddress - defaultShippingAddress"))
}

// AccountType is the resolver for the accountType field.
func (r *accountResolver) AccountType(ctx context.Context, obj *main.Account) (main.AccountType, error) {
	panic(fmt.Errorf("not implemented: AccountType - accountType"))
}

// CreatedAt is the resolver for the createdAt field.
func (r *accountResolver) CreatedAt(ctx context.Context, obj *main.Account) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - createdAt"))
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *accountResolver) UpdatedAt(ctx context.Context, obj *main.Account) (string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updatedAt"))
}

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input main.CreateAccountInput) (*main.Account, error) {
	panic(fmt.Errorf("not implemented: CreateAccount - createAccount"))
}

// UpdateAccount is the resolver for the updateAccount field.
func (r *mutationResolver) UpdateAccount(ctx context.Context, id string, input main.CreateAccountInput) (*main.Account, error) {
	panic(fmt.Errorf("not implemented: UpdateAccount - updateAccount"))
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input main.CreateProductInput) (*main.Product, error) {
	panic(fmt.Errorf("not implemented: CreateProduct - createProduct"))
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, input main.CreateProductInput) (*main.Product, error) {
	panic(fmt.Errorf("not implemented: UpdateProduct - updateProduct"))
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input main.CreateOrderInput) (*main.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// UpdateOrderStatus is the resolver for the updateOrderStatus field.
func (r *mutationResolver) UpdateOrderStatus(ctx context.Context, id string, status main.OrderStatus) (*main.Order, error) {
	panic(fmt.Errorf("not implemented: UpdateOrderStatus - updateOrderStatus"))
}

// CancelOrder is the resolver for the cancelOrder field.
func (r *mutationResolver) CancelOrder(ctx context.Context, id string) (*main.Order, error) {
	panic(fmt.Errorf("not implemented: CancelOrder - cancelOrder"))
}

// CreateShipment is the resolver for the createShipment field.
func (r *mutationResolver) CreateShipment(ctx context.Context, input main.CreateShipmentInput) (*main.Shipment, error) {
	panic(fmt.Errorf("not implemented: CreateShipment - createShipment"))
}

// UpdateShipmentStatus is the resolver for the updateShipmentStatus field.
func (r *mutationResolver) UpdateShipmentStatus(ctx context.Context, id string, status main.ShipmentStatus) (*main.Shipment, error) {
	panic(fmt.Errorf("not implemented: UpdateShipmentStatus - updateShipmentStatus"))
}

// AssignCourierToShipment is the resolver for the assignCourierToShipment field.
func (r *mutationResolver) AssignCourierToShipment(ctx context.Context, shipmentID string, courierID string) (*main.Shipment, error) {
	panic(fmt.Errorf("not implemented: AssignCourierToShipment - assignCourierToShipment"))
}

// Account is the resolver for the account field.
func (r *queryResolver) Account(ctx context.Context, id string) (*main.Account, error) {
	panic(fmt.Errorf("not implemented: Account - account"))
}

// Accounts is the resolver for the accounts field.
func (r *queryResolver) Accounts(ctx context.Context, pagination *main.PaginationInput, accountType *main.AccountType) ([]*main.Account, error) {
	panic(fmt.Errorf("not implemented: Accounts - accounts"))
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*main.Product, error) {
	panic(fmt.Errorf("not implemented: Product - product"))
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, pagination *main.PaginationInput, category *string) ([]*main.Product, error) {
	panic(fmt.Errorf("not implemented: Products - products"))
}

// Order is the resolver for the order field.
func (r *queryResolver) Order(ctx context.Context, id string) (*main.Order, error) {
	panic(fmt.Errorf("not implemented: Order - order"))
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context, pagination *main.PaginationInput, accountID *string, filter *main.OrderFilterInput) ([]*main.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}

// Shipment is the resolver for the shipment field.
func (r *queryResolver) Shipment(ctx context.Context, id string) (*main.Shipment, error) {
	panic(fmt.Errorf("not implemented: Shipment - shipment"))
}

// Shipments is the resolver for the shipments field.
func (r *queryResolver) Shipments(ctx context.Context, pagination *main.PaginationInput, orderID *string, filter *main.ShipmentFilterInput) ([]*main.Shipment, error) {
	panic(fmt.Errorf("not implemented: Shipments - shipments"))
}

// CourierPartner is the resolver for the courierPartner field.
func (r *queryResolver) CourierPartner(ctx context.Context, id string) (*main.CourierPartner, error) {
	panic(fmt.Errorf("not implemented: CourierPartner - courierPartner"))
}

// CourierPartners is the resolver for the courierPartners field.
func (r *queryResolver) CourierPartners(ctx context.Context, pagination *main.PaginationInput, supportedMethod *main.ShippingMethod) ([]*main.CourierPartner, error) {
	panic(fmt.Errorf("not implemented: CourierPartners - courierPartners"))
}

// Account returns main.AccountResolver implementation.
func (r *Resolver) Account() main.AccountResolver { return &accountResolver{r} }

// Mutation returns main.MutationResolver implementation.
func (r *Resolver) Mutation() main.MutationResolver { return &mutationResolver{r} }

// Query returns main.QueryResolver implementation.
func (r *Resolver) Query() main.QueryResolver { return &queryResolver{r} }

type accountResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }