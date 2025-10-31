/*
 * This file is part of YukkiMusic.
 *
 * YukkiMusic — A Telegram bot that streams music into group voice chats with seamless playback and control.
 * Copyright (C) 2025 TheTeamVivek
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <https://www.gnu.org/licenses/>.
 */
package core

import (
	"fmt"

	tg "github.com/amarnathcjd/gogram/telegram"

	"github.com/TheTeamVivek/YukkiMusic/config"
	"github.com/TheTeamVivek/YukkiMusic/internal/utils"
)

func AddMeMarkup(username string) tg.ReplyMarkup {
	return tg.NewKeyboard().
		AddRow(
			tg.Button.URL("♕︎ 𝐀ᴅᴅ 𝐌ᴇ 𝐘ᴏᴜʀ 𝐆ʀᴏᴜᴘ ♕︎",
				"https://t.me/"+username+"?startgroup=true",
			),
		).
		Build()
}

func GetPlayMarkup(r *RoomState, queued bool) tg.ReplyMarkup {
	btn := tg.NewKeyboard()
	prefix := "room:"
	if r.IsCPlay() {
		prefix = "croom:"
	}
	progress := utils.GetProgressBar(r.Position, r.Track.Duration)
	progress = formatDuration(r.Position) + " " + progress + " " + formatDuration(r.Track.Duration)

	if !queued {
		btn.AddRow(
			tg.Button.Data(progress, "progress"),
		)
	}
	btn.AddRow(
		tg.Button.Data("▷", prefix+"resume"),
		tg.Button.Data("II", prefix+"pause"),
		tg.Button.Data("‣‣I", prefix+"skip"),
		tg.Button.Data("▢", prefix+"stop"),
	)

	btn.AddRow(
		tg.Button.Data("↩ 15s", "room:seekback_15"),
		tg.Button.Data("⟳", "room:replay"),
		tg.Button.Data("15s ↪", "room:seek_15"),
	)

	btn.AddRow(
		tg.Button.Data("Close", "close"),
	)

	return btn.Build()
}

func GetGroupHelpKeyboard() *tg.ReplyInlineMarkup {
	return tg.NewKeyboard().
		AddRow(
			tg.Button.URL("📒 𝐂ᴏᴍᴍᴀɴᴅs", "https://t.me/"+BUser.Username+"?start=help"),
		).
		Build()
}

func GetStartMarkup() tg.ReplyMarkup {
	return tg.NewKeyboard().
		AddRow(
			tg.Button.URL("♕︎ 𝐀ᴅᴅ 𝐌ᴇ 𝐘ᴏᴜʀ 𝐆ʀᴏᴜᴘ ♕︎",
				"https://t.me/"+BUser.Username+"?startgroup=true",
			),
		).
		AddRow(
			tg.Button.Data("𝐇ᴇʟᴘ & 𝐂ᴏᴍᴍᴀɴᴅs", "help_cb"),
		//	tg.Button.URL("💻 Source", config.RepoURL),
		).
		AddRow(
			tg.Button.URL("𝐔ᴘᴅᴀᴛᴇ 💬", config.SupportChannel),
			tg.Button.URL("𝐒ᴜᴘᴘᴏʀᴛ 🚀", config.SupportChat),
		).
		Build()
}

func GetHelpKeyboard() *tg.ReplyInlineMarkup {
	return tg.NewKeyboard().
		AddRow(
			tg.Button.Data("🛠 𝐀ᴅᴍɪɴs", "help:admins"),
			tg.Button.Data("🌍 𝐏ᴜʙʟɪᴄ", "help:public"),
		).
		AddRow(
			tg.Button.Data("👑 𝐎ᴡɴᴇʀ", "help:owner"),
			tg.Button.Data("⚡ 𝐒ᴜᴅᴏᴇʀs", "help:sudoers"),
		).
		AddRow(tg.Button.Data("⬅️ 𝐁ᴀᴄᴋ", "start")).
		Build()
}

func GetBackKeyboard() *tg.ReplyInlineMarkup {
	return tg.NewKeyboard().
		AddRow(tg.Button.Data("⬅️ 𝐁ᴀᴄᴋ", "help:main")).
		Build()
}

func formatDuration(sec int) string {
	h := sec / 3600
	m := (sec % 3600) / 60
	s := sec % 60

	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s) // HH:MM:SS
	}
	return fmt.Sprintf("%02d:%02d", m, s) // MM:SS
}
