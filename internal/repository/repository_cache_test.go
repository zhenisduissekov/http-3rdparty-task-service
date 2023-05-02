package repository
import (
"testing"

"github.com/stretchr/testify/assert"
"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
)

func TestRepoCache_Set(t *testing.T) {
	cnf := config.New()
	repoCache := NewCache(cnf)
	repoCache.Set("key", entity.Task{
		Id:          "id",
	})
	id, err := repoCache.Get("key")
	assert.NoError(t, err)
	assert.Equal(t, "id", id.Id)
}


func TestRepoCache_Set_Overwrite(t *testing.T) {
	cnf := config.New()
	repoCache := NewCache(cnf)
	repoCache.Set("key", entity.Task{
		Id:          "id",
	})
	
	repoCache.Set("key", entity.Task{
		Id:          "id2",
	})
	id, err := repoCache.Get("key")
	assert.NoError(t, err)
	assert.Equal(t, "id2", id.Id)
}


func TestRepoCache_Set_Not_Found(t *testing.T) {
	cnf := config.New()
	repoCache := NewCache(cnf)
	id, err := repoCache.Get("key")
	assert.Error(t, err)
	assert.Equal(t, "no task found", err.Error())
	assert.Equal(t, id, entity.Task{})
}