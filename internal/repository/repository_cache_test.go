package repository
import (
"testing"

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
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if id.Id != "id" {
		t.Errorf("error: %v", err)
	}
}
