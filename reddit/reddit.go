package reddit

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type apiResponse struct {
    Kind string `json:"kind"`
    Data apiResponseData `json:"data"`
}

type apiResponseData struct {
    Modhash string `json:"modhash"`
    Children []apiChild `json:"children"`
}

type apiChild struct {
    Kind string `json:"kind"`
    Data Submission `json:"data"`
}

type Submission struct {
    Title string `json:"title"`
    Domain string `json:"domain"`
    Url string `json:"url"`
    Author string `json:"author"`
    Score int `json:"score"`
    Selftext_html *string `json:"selftext_html"`
    Subreddit string `json:"subreddit"`
    Id string `json:"id"`
    Clicked bool `json:"clicked"`
    Over_18 bool `json:"over_18"`
    Hidden bool `json:"hidden"`
    Saved bool `json:"saved"`
    Is_self bool `json:"is_self"`
    Edited interface{} `json:"edited"` // bool or float64. reddit sucks
    Thumbnail string `json:"thumbnail"`
    Subreddit_id string `json:"subreddit_id"`
    Downs int `json:"downs"`
    Permalink string `json:"permalink"`
    Name string `json:"name"`
    Created float64 `json:"created"`
    Created_utc float64 `json:"created_utc"`
    Ups int `json:"ups"`
    Num_comments int `json:"num_comments"`
    // num_reports ?
    // distinguished ?
    // banned_by ?
    // media_embed ?
    // author_flair_text ?
    // link_flair_css_class ?
    // author_flair_css_class ?
    // media ?
    // approved_by ?
    // likes ?
    // link_flair_text ?
}

func GetSubmissions(subreddit string) ([]Submission, error) {
    url := fmt.Sprintf("http://reddit.com/r/%s.json", subreddit)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    var apiResp apiResponse
    err = json.Unmarshal(body, &apiResp)
    if err != nil {
        return nil, err
    }
    ret := make([]Submission, len(apiResp.Data.Children))
    for i, v := range apiResp.Data.Children {
        ret[i] = v.Data
    }

    return ret, nil
}
