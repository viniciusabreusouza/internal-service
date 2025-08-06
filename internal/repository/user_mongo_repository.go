package repository

import (
	"context"
	"errors"
	"time"

	"example.com/internal-service/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CollectionInterface define a interface para operações de collection
type CollectionInterface interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
}

// UserMongoRepository implementa UserRepository usando MongoDB
type UserMongoRepository struct {
	collection CollectionInterface
}

// NewUserMongoRepository cria uma nova instância do repositório MongoDB
func NewUserMongoRepository(db *mongo.Database) *UserMongoRepository {
	return &UserMongoRepository{
		collection: db.Collection("users"),
	}
}

// NewUserMongoRepositoryWithCollection cria uma nova instância do repositório MongoDB com uma collection customizada
func NewUserMongoRepositoryWithCollection(collection CollectionInterface) *UserMongoRepository {
	return &UserMongoRepository{
		collection: collection,
	}
}

// userDocument representa o documento MongoDB para User
type userDocument struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"user_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// toDocument converte User para userDocument
func toDocument(u *user.User) userDocument {
	return userDocument{
		UserID:    u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// toUser converte userDocument para User
func (d userDocument) toUser() *user.User {
	return &user.User{
		ID:        d.UserID,
		Name:      d.Name,
		Email:     d.Email,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// Create cria um novo usuário no MongoDB
func (r *UserMongoRepository) Create(ctx context.Context, user *user.User) error {
	doc := toDocument(user)

	_, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}

// GetByID busca um usuário pelo ID
func (r *UserMongoRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	var doc userDocument

	err := r.collection.FindOne(ctx, bson.M{"user_id": id}).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return doc.toUser(), nil
}

// GetAll busca todos os usuários com paginação
func (r *UserMongoRepository) GetAll(ctx context.Context, page, limit int) ([]*user.User, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	skip := (page - 1) * limit

	// Conta o total de documentos
	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	// Busca os documentos com paginação
	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var users []*user.User
	for cursor.Next(ctx) {
		var doc userDocument
		if err := cursor.Decode(&doc); err != nil {
			return nil, 0, err
		}
		users = append(users, doc.toUser())
	}

	if err := cursor.Err(); err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

// Update atualiza um usuário existente
func (r *UserMongoRepository) Update(ctx context.Context, user *user.User) error {
	user.UpdatedAt = time.Now()
	doc := toDocument(user)

	filter := bson.M{"user_id": user.ID}
	update := bson.M{
		"$set": bson.M{
			"name":       doc.Name,
			"email":      doc.Email,
			"updated_at": doc.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// Delete remove um usuário pelo ID
func (r *UserMongoRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"user_id": id}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}
