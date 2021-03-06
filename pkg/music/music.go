package music

import (
	"bufio"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"layeh.com/gopus"
)

const (
	YOUTUBE_EXTRACTOR = "youtube:playlist"
)

type PlaylistSong struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Type  string `json:"_type"`
	IeKey string `json:"ie_key"`
	Title string `json:"title"`
}

type Song struct {
	EndTime            interface{}            `json:"end_time"`
	UploaderURL        string                 `json:"uploader_url"`
	ViewCount          interface{}            `json:"view_count"`
	DislikeCount       int                    `json:"dislike_count"`
	Format             string                 `json:"format"`
	Categories         []string               `json:"categories"`
	Height             int                    `json:"height"`
	Ext                string                 `json:"ext"`
	UploaderID         string                 `json:"uploader_id"`
	Formats            []SongFormats          `json:"formats"`
	Acodec             string                 `json:"acodec"`
	Subtitles          SongSubtitles          `json:"subtitles"`
	AutomaticCaptions  SongAutomaticCaptions  `json:"automatic_captions"`
	AgeLimit           int                    `json:"age_limit"`
	Uploader           string                 `json:"uploader"`
	ExtractorKey       string                 `json:"extractor_key"`
	Annotations        interface{}            `json:"annotations"`
	RequestedFormats   []SongRequestedFormats `json:"requested_formats"`
	EpisodeNumber      interface{}            `json:"episode_number"`
	Fps                float64                `json:"fps"`
	FormatID           string                 `json:"format_id"`
	Series             interface{}            `json:"series"`
	StretchedRatio     interface{}            `json:"stretched_ratio"`
	DisplayID          string                 `json:"display_id"`
	LikeCount          int                    `json:"like_count"`
	Tags               []string               `json:"tags"`
	IsLive             interface{}            `json:"is_live"`
	Creator            interface{}            `json:"creator"`
	WebpageURL         string                 `json:"webpage_url"`
	Resolution         interface{}            `json:"resolution"`
	Description        string                 `json:"description"`
	UploadDate         string                 `json:"upload_date"`
	Chapters           []SongChapters         `json:"chapters"`
	ID                 string                 `json:"id"`
	Width              int                    `json:"width"`
	Vcodec             string                 `json:"vcodec"`
	PlaylistIndex      interface{}            `json:"playlist_index"`
	AltTitle           interface{}            `json:"alt_title"`
	License            interface{}            `json:"license"`
	Abr                float64                `json:"abr"`
	Extractor          string                 `json:"extractor"`
	Duration           int                    `json:"duration"`
	StartTime          interface{}            `json:"start_time"`
	Thumbnail          string                 `json:"thumbnail"`
	Vbr                interface{}            `json:"vbr"`
	SeasonNumber       interface{}            `json:"season_number"`
	Title              string                 `json:"title"`
	RequestedSubtitles interface{}            `json:"requested_subtitles"`
	WebpageURLBasename string                 `json:"webpage_url_basename"`
	Track              interface{}            `json:"track"`
	Artist             interface{}            `json:"artist"`
	Album              interface{}            `json:"album"`
	Thumbnails         []SongThumbnails       `json:"thumbnails"`
	AverageRating      interface{}            `json:"average_rating"`
	Playlist           interface{}            `json:"playlist"`
}
type SongDownloaderOptions struct {
	HTTPChunkSize int `json:"http_chunk_size"`
}
type SongHTTPHeaders struct {
	Accept         string `json:"Accept"`
	AcceptEncoding string `json:"Accept-Encoding"`
	AcceptLanguage string `json:"Accept-Language"`
	UserAgent      string `json:"User-Agent"`
	AcceptCharset  string `json:"Accept-Charset"`
}
type SongFormats struct {
	Fps               interface{}           `json:"fps"`
	Abr               float64               `json:"abr,omitempty"`
	Quality           int                   `json:"quality"`
	FormatNote        string                `json:"format_note"`
	Format            string                `json:"format"`
	DownloaderOptions SongDownloaderOptions `json:"downloader_options,omitempty"`
	Vcodec            string                `json:"vcodec"`
	Acodec            string                `json:"acodec"`
	URL               string                `json:"url"`
	Protocol          string                `json:"protocol"`
	Filesize          int                   `json:"filesize"`
	Asr               int                   `json:"asr"`
	HTTPHeaders       SongHTTPHeaders       `json:"http_headers"`
	Container         string                `json:"container,omitempty"`
	FormatID          string                `json:"format_id"`
	Ext               string                `json:"ext"`
	Height            interface{}           `json:"height"`
	Tbr               float64               `json:"tbr"`
	Width             interface{}           `json:"width"`
	Vbr               float64               `json:"vbr,omitempty"`
}
type SongSubtitles struct {
}
type SongAutomaticCaptions struct {
}
type SongRequestedFormats struct {
	Fps               int                   `json:"fps"`
	Quality           int                   `json:"quality"`
	FormatNote        string                `json:"format_note"`
	Vcodec            string                `json:"vcodec"`
	DownloaderOptions SongDownloaderOptions `json:"downloader_options"`
	Format            string                `json:"format"`
	Acodec            string                `json:"acodec"`
	Vbr               float64               `json:"vbr,omitempty"`
	URL               string                `json:"url"`
	Protocol          string                `json:"protocol"`
	Filesize          int                   `json:"filesize"`
	Asr               interface{}           `json:"asr"`
	HTTPHeaders       SongHTTPHeaders       `json:"http_headers"`
	Container         string                `json:"container"`
	FormatID          string                `json:"format_id"`
	Ext               string                `json:"ext"`
	Height            int                   `json:"height"`
	Tbr               float64               `json:"tbr"`
	Width             int                   `json:"width"`
	Abr               float64               `json:"abr,omitempty"`
}
type SongChapters struct {
	EndTime   float64 `json:"end_time"`
	Title     string  `json:"title"`
	StartTime float64 `json:"start_time"`
}
type SongThumbnails struct {
	URL string `json:"url"`
	ID  string `json:"id"`
}

type Playlist struct {
	Type               string            `json:"_type"`
	WebpageURLBasename string            `json:"webpage_url_basename"`
	ExtractorKey       string            `json:"extractor_key"`
	ID                 string            `json:"id"`
	WebpageURL         string            `json:"webpage_url"`
	Extractor          string            `json:"extractor"`
	Title              string            `json:"title"`
	Entries            []PlaylistEntries `json:"entries"`
}
type PlaylistEntries struct {
	Type  string `json:"_type"`
	IeKey string `json:"ie_key"`
	ID    string `json:"id"`
	URL   string `json:"url"`
}

// handleSongProcessCleanup will wait for a process to end so it can be cleaned up. Will cancel error checking if
// passed context is cancelled. Will log an error with your message if there's an error without context being
// cancelled.
func handleSongProcessCleanup(ctx context.Context, c *exec.Cmd, log *logrus.Logger, errorMessage string) {
	select {
	case <-ctx.Done():
		// Song was skipped or we're shutting down so the exit code will be 1 and return an error no matter what.
		_ = c.Wait()
	default:
		err := c.Wait()
		if err != nil {
			log.WithError(err).Error(errorMessage)
		}
	}
}

func StreamSong(ctx context.Context, link string, log *logrus.Logger, vc *discordgo.VoiceConnection, volume float32) {
	// TODO remove log.fatal
	cmd := exec.CommandContext(ctx, "youtube-dl", "--no-progress", "--no-call-home", "--default-search", "ytsearch", "--no-playlist", "--no-mtime", "-o", "-", "--format", "bestaudio/worstvideo/best", "--prefer-ffmpeg", "--quiet", link)
	run := exec.CommandContext(ctx, "ffmpeg", "-i", "-", "-vn", "-acodec", "pcm_s16le", "-f", "s16le", "-ar", "48000", "-af", fmt.Sprintf("volume=%f", volume), "-ac", "2", "pipe:1")
	ytdl, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	r := bufio.NewReaderSize(ytdl, 16384*8)
	run.Stdin = r
	ffmpegout, err := run.StdoutPipe()
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	// Cleanup process after the process ends.
	defer handleSongProcessCleanup(ctx, cmd, log, "Unable to close ffmpeg process.")

	// run.Stderr = os.Stderr
	err = run.Start()
	if err != nil {
		log.Fatal(err)
	}
	// Cleanup process after the process ends.
	defer handleSongProcessCleanup(ctx, run, log, "Unable to close ffmpeg process.")

	ffmpegbuf := bufio.NewReaderSize(ffmpegout, 16384)
	err = vc.Speaking(true)
	if err != nil {
		log.WithError(err).Error("Unable to set voice connection to speaking.")
	}
	defer func() {
		err := vc.Speaking(false)
		if err != nil {
			log.WithError(err).Error("Unable to set voice connection to not speaking.")
		}
	}()
	opusEncoder, err := gopus.NewEncoder(48000, 2, gopus.Audio)

	if err != nil {
		fmt.Println("NewEncoder Error:", err)
		return
	}
	for {
		// Increase audiobuf for fun.
		audiobuf := make([]int16, 960*2)
		err = binary.Read(ffmpegbuf, binary.LittleEndian, &audiobuf)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			// fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Println("error reading from ffmpeg stdout :", err)
			break
		}
		select {
		default:
			select {
			case <-ctx.Done():
				return
			default:
				opus, err := opusEncoder.Encode(audiobuf, 960, 3840)
				if err != nil {
					fmt.Println("Encoding Error:", err)
					return
				}
				if vc.OpusSend == nil {
					fmt.Printf("Discordgo not ready for opus packets.\n")
					return
				}
				vc.OpusSend <- opus
			}
		}
	}
}

func GetSongInfo(ctx context.Context, url string) (*Song, error) {
	song := &Song{}
	cmd := exec.CommandContext(ctx, "youtube-dl", "--simulate", "--print-json", "--no-progress", "--no-call-home", "--default-search", "ytsearch", "--no-playlist", "--no-mtime", url)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return song, err
	}
	err = json.Unmarshal(output, song)
	if err != nil {
		return song, err
	}
	return song, nil
}

func GetPlaylistInfo(ctx context.Context, url string) ([]*PlaylistSong, error) {
	cmd := exec.CommandContext(ctx, "youtube-dl", "--simulate", "--dump-json", "--no-progress", "--no-call-home", "--default-search", "ytsearch", "--flat-playlist", "--no-mtime", url)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	var playlistSongs []*PlaylistSong
	outputString := string(output[:])
	log.Println(outputString)
	songSplit := strings.Split(outputString, "\n")
	for _, songString := range songSplit {
		if songString != "" {
			song := &PlaylistSong{}
			err := json.Unmarshal([]byte(songString), song)
			if err == nil {
				if song.Type == "url" {
					switch song.IeKey {
					case "Youtube":
						song.URL = fmt.Sprintf("https://youtu.be/%s", song.URL)
					}
				}
				playlistSongs = append(playlistSongs, song)
			}
		}
	}
	return playlistSongs, nil
}
