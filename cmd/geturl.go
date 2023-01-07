package cmd

// func getUrlFile(data []entities.ListDownload) string {
// 	conf := config.ReadConfig()
// 	var result string
// 	for _, d := range data {
// 		if strings.Contains(d.Codec, conf.Codec) {
// 			for _, urlfile := range d.Downloads {
// 				if strings.Contains(urlfile.Resolution, conf.Resolution) {
// 					for _, u := range urlfile.Links {
// 						if strings.Contains(u.Label, conf.FileHosting) {
// 							result = u.Link
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return result
// }
//
// func getResponse(u *url.URL) scrape.Response {
// 	var result scrape.Response
// 	switch {
// 	case strings.Contains(u.Host, "15.235.11.45"):
// 		result = scrape.Oploverz(u.String())
// 	case strings.Contains(u.Host, "doronime"):
// 		result = scrape.Doronime(u.String())
// 	case strings.Contains(u.Host, "samehadaku.win"):
// 		result = scrape.Samehadaku(u.String())
// 	case strings.Contains(u.Host, "lendrive"):
// 		result = scrape.Lendrive(u.String())
// 	}
// 	return result
// }
//
// var getUrlCmd = &cobra.Command{
// 	Use:   "get",
// 	Short: "Get url download",
// 	Long:  "Get url download from spesific website",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		if len(args) <= 0 || len(args) >= 2 {
// 			cmd.Help()
// 			return
// 		}
//
// 		urlWeb := args[0]
// 		conf := config.ReadConfig()
//
// 		u, err := url.Parse(urlWeb)
// 		if err != nil {
// 			fmt.Println("Url Tidak Valid")
// 			return
// 		}
//
// 		response := getResponse(u)
//
// 		raw, _ := cmd.Flags().GetBool("raw")
// 		if raw {
// 			rawJson, _ := json.Marshal(response)
// 			fmt.Println(string(rawJson))
// 			return
// 		}
//
// 		result := getUrlFile(response.Downloads)
//
// 		fmt.Println("[ Open URL ] >> ", urlWeb)
// 		fmt.Println("[ Result   ] >> ", result)
//
// 		if result != "" && conf.OpenInBrowser == "true" {
// 			if exec.Command(conf.Browser, result).Run(); err != nil {
// 				fmt.Println(err.Error())
// 				return
// 			}
// 		}
//
// 	},
// }
//
// func init() {
// 	getUrlCmd.PersistentFlags().Bool("raw", false, "Print raw response in json")
// 	RootCmd.AddCommand(getUrlCmd)
// }
