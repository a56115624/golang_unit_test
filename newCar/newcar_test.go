package newcar

import (
	"testing"
	"time"
)

func TestnewCar(t *testing.T) {

for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		e := 
		
		}
		if got := e.CalculateMoney(tt.args.money); got != tt.want {
			t.Errorf("CalculateMoney() = %v, want %v", got, tt.want)
		}
	})
}
	

package session_handler

import (
	"fmt"
	"github.com/golang/mock/gomock"
	customError "github.com/teampui/comico-api/pkg/custom-error"
	"github.com/teampui/comico-api/pkg/model"
	track "github.com/teampui/tracking"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestContentHandler_handleItem(t *testing.T) {
	initContentHandler(t)

	mockDomainService.EXPECT().GetResourceDomain().Return("src.com", nil).AnyTimes()

	app.Get("/api/vi/contents/items/:item_id", contentHandler.handleItem)
	tests := []struct {
		name   string
		route  string
		want   any
		invoke func(id string)
	}{
		{"handleItem Success", "/api/vi/contents/items/1",
			`{"data":{"id":1,"title":"test","creator":"test","description":"test","status":"連載","covers":{"thumb":"test.js"},"hot":"1","views":"11","bookmark_status":false,"chapter":null},"domain":"src.com","status":200}`,
			func(id string) {
				mockContentService.EXPECT().GetItemByItemId(false, int64(0), "1").Return(&model.ItemDetail{
					ItemInfo: model.ItemInfo{
						Id:          1,
						Title:       "test",
						Creator:     "test",
						Description: "test",
						Status:      "連載",
						Covers:      map[string]any{"thumb": "test.js"},
						Hot:         "1",
						Views:       "11",
					},
					BookmarkStatus: false,
					Chapter:        nil,
				}, nil)

				mockTrackingService.EXPECT().InsertLog(&track.Tracking{
					Referrer: fmt.Sprintf("%d", 0),
					Platform: "web",
					Event:    "item_view",
					Object:   "1",
					Uid2:     fmt.Sprintf("%d", 0),
					IP:       "0.0.0.0",
					Version:  "1",
				}).Return(nil)
			}},
		{"handleItem Success", "/api/vi/contents/items/2",
			`{"data":{"id":2,"title":"test","creator":"test","description":"test","status":"連載","covers":{"thumb":"test.js"},"hot":"1","views":"11","bookmark_status":false,"chapter":null},"domain":"src.com","status":200}`,
			func(id string) {
				mockContentService.EXPECT().GetItemByItemId(false, int64(0), "2").Return(&model.ItemDetail{
					ItemInfo: model.ItemInfo{
						Id:          2,
						Title:       "test",
						Creator:     "test",
						Description: "test",
						Status:      "連載",
						Covers:      map[string]any{"thumb": "test.js"},
						Hot:         "1",
						Views:       "11",
					},
					BookmarkStatus: false,
					Chapter:        nil,
				}, nil)

				mockTrackingService.EXPECT().InsertLog(&track.Tracking{
					Referrer: fmt.Sprintf("%d", 0),
					Platform: "web",
					Event:    "item_view",
					Object:   "2",
					Uid2:     fmt.Sprintf("%d", 0),
					IP:       "0.0.0.0",
					Version:  "1",
				}).Return(nil)
			},
		},
		{"handleItem failed", "/api/vi/contents/items/2",
			`{"status":500,"message":"sql error","error":true}`,
			func(id string) {
				mockContentService.EXPECT().GetItemByItemId(false, int64(0), "2").
					Return(nil, &customError.ComicoError{
						Status:    500,
						Message:   "sql error",
						ErrorBool: true,
						ErrorMsg:  "sql error",
					})

				mockLogService.EXPECT().AddLog(gomock.Any()).Return(nil)

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.invoke(sessionId)
			req := httptest.NewRequest("GET", tt.route, nil)
			resp, _ := app.Test(req, -1)

			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()

			if !reflect.DeepEqual(string(body), tt.want) {
				t.Errorf("ContentHandler.handleItem() = %v, want %v", string(body), tt.want)
			}
		})
	}
}