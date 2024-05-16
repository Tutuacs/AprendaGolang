package db

import "context"

func Insert(collection string, data any) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	col := client.Database("webform").Collection(collection)

	_, err := col.InsertOne(context.Background(), data)

	return err
}
