package models

type Media struct {
	LogoImagePath          string `json:"logo_image_path"`
	ChannelBannerImagePath string `json:"channel_banner_image_path"`
}
type ChannelModel struct {
	ChannelID               int    `json:"channel_id"`
	ChannelName             string `json:"channel_name"`
	ChannelRank             int    `json:"channel_rank"`
	ChannelSubscribersCount int    `json:"channel_subscribers_count"`
	ChannelViewsCount       int    `json:"channel_views_count"`
	ChannelVideosCount      int    `json:"channel_videos_count"`
	Media                   Media  `json:"media"`
}

type ChannelsModel struct {
	Channels []ChannelModel `json:"channels"`
}
