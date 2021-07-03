package dao

import (
	"fmt"
	"github.com/my/repo/src/database"
	"github.com/my/repo/src/models"
)

func GetAllChannelsDao(sortBy string, orderBy string, pageNo int64) (channels models.ChannelsModel, err error) {
	var query = fmt.Sprint(
		`select channel_id,channel_name,channel_rank,channel_subscribers_count,channel_videos_count,channel_views_count,media::json from channels order by `,
		ChannelOrderByGenerator(orderBy, sortBy), GetLimitAndOffset(pageNo))
	_, err = database.DBInstance.Query(&channels.Channels, query)
	if channels.Channels == nil || len(channels.Channels) == 0 {
		var channel models.ChannelsModel
		channels.Channels = channel.Channels
	}
	return channels, err
}

func ChannelOrderByGenerator(sortOrder string, columnName string) string {
	if len(sortOrder) > 0 {
		if len(columnName) > 0 {
			return fmt.Sprint(columnName, ` `, sortOrder)
		}
	}
	return "channel_name asc"
}

func GetLimitAndOffset(pageNo int64) string {
	if pageNo < 2 {
		return fmt.Sprint(` limit `, 10, `offset `, 0)
	}
	return fmt.Sprint(` limit `, 10, ` offset `, pageNo*10)
}
