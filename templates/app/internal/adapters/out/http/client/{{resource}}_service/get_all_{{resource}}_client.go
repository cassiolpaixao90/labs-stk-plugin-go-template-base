package {{resource}}_service
 
import (
	"context"
	"encoding/json"
	"io"
	"{{project_name}}/internal/adapter/out/http/client/{{resource}}_service/dtos"
	"{{project_name}}/internal/adapter/out/http/helpers"
	"{{project_name}}/internal/application/domain"
	"{{project_name}}/internal/application/ports/out"
	"net/http"
)

type GetAll{{resource|capitalize}}Client struct {
	client helpers.HttpClientRequest
}

func NewGetAll{{resource|capitalize}}Client() out.GetAll{{resource|capitalize}}ClientPort {
	return &GetAll{{resource|capitalize}}Client{}
}

func (g *GetAll{{resource|capitalize}}Client) GetAll{{resource|capitalize}}(ctx context.Context, id string) (*domain.{{resource|capitalize}}Domain, error) {

	req, _ := http.NewRequest(http.MethodGet, "url", nil)

	err, resp := g.client.Retry(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// TODO add logger
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	var response *dtos.GetAll{{resource|capitalize}}Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.ToDomain(), nil
}
