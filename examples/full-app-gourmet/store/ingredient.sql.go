// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: ingredient.sql

package store

import (
	"context"

	"simple-crud/store/types"
)

const createIngredient = `-- name: CreateIngredient :one
INSERT INTO ingredient 
(id, name, description, available_all_year, available_jan, available_feb, available_mar, available_apr, available_may, available_jun, available_jul, available_aug, available_sep, available_oct, available_nov, available_dec, category, default_unit)
VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
RETURNING id, created_at, name, description, default_unit, category, available_all_year, available_jan, available_feb, available_mar, available_apr, available_may, available_jun, available_jul, available_aug, available_sep, available_oct, available_nov, available_dec
`

type CreateIngredientParams struct {
	ID               string         `json:"id"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	AvailableAllYear bool           `json:"available_all_year"`
	AvailableJan     bool           `json:"available_jan"`
	AvailableFeb     bool           `json:"available_feb"`
	AvailableMar     bool           `json:"available_mar"`
	AvailableApr     bool           `json:"available_apr"`
	AvailableMay     bool           `json:"available_may"`
	AvailableJun     bool           `json:"available_jun"`
	AvailableJul     bool           `json:"available_jul"`
	AvailableAug     bool           `json:"available_aug"`
	AvailableSep     bool           `json:"available_sep"`
	AvailableOct     bool           `json:"available_oct"`
	AvailableNov     bool           `json:"available_nov"`
	AvailableDec     bool           `json:"available_dec"`
	Category         types.Category `json:"category"`
	DefaultUnit      types.Unit     `json:"default_unit"`
}

func (q *Queries) CreateIngredient(ctx context.Context, arg CreateIngredientParams) (Ingredient, error) {
	row := q.db.QueryRowContext(ctx, createIngredient,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.AvailableAllYear,
		arg.AvailableJan,
		arg.AvailableFeb,
		arg.AvailableMar,
		arg.AvailableApr,
		arg.AvailableMay,
		arg.AvailableJun,
		arg.AvailableJul,
		arg.AvailableAug,
		arg.AvailableSep,
		arg.AvailableOct,
		arg.AvailableNov,
		arg.AvailableDec,
		arg.Category,
		arg.DefaultUnit,
	)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.DefaultUnit,
		&i.Category,
		&i.AvailableAllYear,
		&i.AvailableJan,
		&i.AvailableFeb,
		&i.AvailableMar,
		&i.AvailableApr,
		&i.AvailableMay,
		&i.AvailableJun,
		&i.AvailableJul,
		&i.AvailableAug,
		&i.AvailableSep,
		&i.AvailableOct,
		&i.AvailableNov,
		&i.AvailableDec,
	)
	return i, err
}

const getIngredient = `-- name: GetIngredient :one
SELECT id, created_at, name, description, default_unit, category, available_all_year, available_jan, available_feb, available_mar, available_apr, available_may, available_jun, available_jul, available_aug, available_sep, available_oct, available_nov, available_dec FROM ingredient WHERE id = ?
`

func (q *Queries) GetIngredient(ctx context.Context, id string) (Ingredient, error) {
	row := q.db.QueryRowContext(ctx, getIngredient, id)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.DefaultUnit,
		&i.Category,
		&i.AvailableAllYear,
		&i.AvailableJan,
		&i.AvailableFeb,
		&i.AvailableMar,
		&i.AvailableApr,
		&i.AvailableMay,
		&i.AvailableJun,
		&i.AvailableJul,
		&i.AvailableAug,
		&i.AvailableSep,
		&i.AvailableOct,
		&i.AvailableNov,
		&i.AvailableDec,
	)
	return i, err
}

const getIngredients = `-- name: GetIngredients :many
SELECT id, created_at, name, description, default_unit, category, available_all_year, available_jan, available_feb, available_mar, available_apr, available_may, available_jun, available_jul, available_aug, available_sep, available_oct, available_nov, available_dec FROM ingredient
`

func (q *Queries) GetIngredients(ctx context.Context) ([]Ingredient, error) {
	rows, err := q.db.QueryContext(ctx, getIngredients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ingredient
	for rows.Next() {
		var i Ingredient
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Description,
			&i.DefaultUnit,
			&i.Category,
			&i.AvailableAllYear,
			&i.AvailableJan,
			&i.AvailableFeb,
			&i.AvailableMar,
			&i.AvailableApr,
			&i.AvailableMay,
			&i.AvailableJun,
			&i.AvailableJul,
			&i.AvailableAug,
			&i.AvailableSep,
			&i.AvailableOct,
			&i.AvailableNov,
			&i.AvailableDec,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getIngredientsOfRecipe = `-- name: GetIngredientsOfRecipe :many
SELECT quantity, unit, ingredient.id, ingredient.created_at, ingredient.name, ingredient.description, ingredient.default_unit, ingredient.category, ingredient.available_all_year, ingredient.available_jan, ingredient.available_feb, ingredient.available_mar, ingredient.available_apr, ingredient.available_may, ingredient.available_jun, ingredient.available_jul, ingredient.available_aug, ingredient.available_sep, ingredient.available_oct, ingredient.available_nov, ingredient.available_dec FROM ingredient
JOIN dosing ON ingredient.id = dosing.ingredient_id
WHERE dosing.recipe_id = ?
`

type GetIngredientsOfRecipeRow struct {
	Quantity   int64      `json:"quantity" validate:"required,gt=0"`
	Unit       types.Unit `json:"unit" validate:"required"`
	Ingredient Ingredient `json:"ingredient"`
}

func (q *Queries) GetIngredientsOfRecipe(ctx context.Context, recipeID string) ([]GetIngredientsOfRecipeRow, error) {
	rows, err := q.db.QueryContext(ctx, getIngredientsOfRecipe, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetIngredientsOfRecipeRow
	for rows.Next() {
		var i GetIngredientsOfRecipeRow
		if err := rows.Scan(
			&i.Quantity,
			&i.Unit,
			&i.Ingredient.ID,
			&i.Ingredient.CreatedAt,
			&i.Ingredient.Name,
			&i.Ingredient.Description,
			&i.Ingredient.DefaultUnit,
			&i.Ingredient.Category,
			&i.Ingredient.AvailableAllYear,
			&i.Ingredient.AvailableJan,
			&i.Ingredient.AvailableFeb,
			&i.Ingredient.AvailableMar,
			&i.Ingredient.AvailableApr,
			&i.Ingredient.AvailableMay,
			&i.Ingredient.AvailableJun,
			&i.Ingredient.AvailableJul,
			&i.Ingredient.AvailableAug,
			&i.Ingredient.AvailableSep,
			&i.Ingredient.AvailableOct,
			&i.Ingredient.AvailableNov,
			&i.Ingredient.AvailableDec,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchIngredients = `-- name: SearchIngredients :many
SELECT id, created_at, name, description, default_unit, category, available_all_year, available_jan, available_feb, available_mar, available_apr, available_may, available_jun, available_jul, available_aug, available_sep, available_oct, available_nov, available_dec FROM ingredient
WHERE name LIKE ?
ORDER BY name ASC
LIMIT ?
OFFSET ?
`

type SearchIngredientsParams struct {
	Name   string `json:"name"`
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
}

func (q *Queries) SearchIngredients(ctx context.Context, arg SearchIngredientsParams) ([]Ingredient, error) {
	rows, err := q.db.QueryContext(ctx, searchIngredients, arg.Name, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ingredient
	for rows.Next() {
		var i Ingredient
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Description,
			&i.DefaultUnit,
			&i.Category,
			&i.AvailableAllYear,
			&i.AvailableJan,
			&i.AvailableFeb,
			&i.AvailableMar,
			&i.AvailableApr,
			&i.AvailableMay,
			&i.AvailableJun,
			&i.AvailableJul,
			&i.AvailableAug,
			&i.AvailableSep,
			&i.AvailableOct,
			&i.AvailableNov,
			&i.AvailableDec,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateIngredient = `-- name: UpdateIngredient :one
UPDATE ingredient SET 
  name=COALESCE(?1, name),
  description=COALESCE(?2, description),
  category=COALESCE(?3, category),
  default_unit=COALESCE(?4, default_unit),
  available_all_year=COALESCE(?5, available_all_year),
  available_jan=COALESCE(?6, available_jan),
  available_feb=COALESCE(?7, available_feb),
  available_mar=COALESCE(?8, available_mar),
  available_apr=COALESCE(?9, available_apr),
  available_may=COALESCE(?10, available_may),
  available_jun=COALESCE(?11, available_jun),
  available_jul=COALESCE(?12, available_jul),
  available_aug=COALESCE(?13, available_aug),
  available_sep=COALESCE(?14, available_sep),
  available_oct=COALESCE(?15, available_oct),
  available_nov=COALESCE(?16, available_nov),
  available_dec=COALESCE(?17, available_dec)
WHERE id = ?18
RETURNING id, created_at, name, description, default_unit, category, available_all_year, available_jan, available_feb, available_mar, available_apr, available_may, available_jun, available_jul, available_aug, available_sep, available_oct, available_nov, available_dec
`

type UpdateIngredientParams struct {
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	Category         types.Category `json:"category"`
	DefaultUnit      types.Unit     `json:"default_unit"`
	AvailableAllYear bool           `json:"available_all_year"`
	AvailableJan     bool           `json:"available_jan"`
	AvailableFeb     bool           `json:"available_feb"`
	AvailableMar     bool           `json:"available_mar"`
	AvailableApr     bool           `json:"available_apr"`
	AvailableMay     bool           `json:"available_may"`
	AvailableJun     bool           `json:"available_jun"`
	AvailableJul     bool           `json:"available_jul"`
	AvailableAug     bool           `json:"available_aug"`
	AvailableSep     bool           `json:"available_sep"`
	AvailableOct     bool           `json:"available_oct"`
	AvailableNov     bool           `json:"available_nov"`
	AvailableDec     bool           `json:"available_dec"`
	ID               string         `json:"id"`
}

func (q *Queries) UpdateIngredient(ctx context.Context, arg UpdateIngredientParams) (Ingredient, error) {
	row := q.db.QueryRowContext(ctx, updateIngredient,
		arg.Name,
		arg.Description,
		arg.Category,
		arg.DefaultUnit,
		arg.AvailableAllYear,
		arg.AvailableJan,
		arg.AvailableFeb,
		arg.AvailableMar,
		arg.AvailableApr,
		arg.AvailableMay,
		arg.AvailableJun,
		arg.AvailableJul,
		arg.AvailableAug,
		arg.AvailableSep,
		arg.AvailableOct,
		arg.AvailableNov,
		arg.AvailableDec,
		arg.ID,
	)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.DefaultUnit,
		&i.Category,
		&i.AvailableAllYear,
		&i.AvailableJan,
		&i.AvailableFeb,
		&i.AvailableMar,
		&i.AvailableApr,
		&i.AvailableMay,
		&i.AvailableJun,
		&i.AvailableJul,
		&i.AvailableAug,
		&i.AvailableSep,
		&i.AvailableOct,
		&i.AvailableNov,
		&i.AvailableDec,
	)
	return i, err
}