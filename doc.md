## [searchMusicTosilk](https://github.com/xiyaowong/searchMusicTosilk/)



1. 首先确保安装好 `FFmpeg` 且位于环境变量

2. 全部都是**GET**请求

   1. 搜索音乐 `/music`

      | param    | 必选 | 可选项                                                       | 说明                       |
      | -------- | ---- | ------------------------------------------------------------ | -------------------------- |
      | `option` | 是   | 1. `qq` -> QQ 音乐 <br />2. `netease` -> 网易云<br />3. `migu` -> 咪咕音乐<br />4. `kugou` ->  酷狗音乐<br />5. `kuwo` ->  酷我音乐 | ip 可能会被咪咕 api ban 掉 |
      | `name`   | 是   | 歌曲名称                                                     |                            |
      | `silk`   | 否   | 任意值，只要存在，则会额外输出转码后 silk 格式的 base64 编码 | 慢                         |

      输出:

      ```json
      {
        "code": 0,
        "msg": "",
        "result": {
          "author": "",
          "song_name": "",
          "song_url": "",
          "song_silk_base64": ""
        }
      }
      ```

      示例: `GET http://localhost:3317/music?option=qq&name=我和我的祖国`

      

   2. 音频链接转码成 silk 格式 ，输出 base64 编码 `/silk`

      | param | 必选 | 可选项   | 说明            |
      | ----- | ---- | -------- | --------------- |
      | `url` | 是   | 音频链接 | 没做任何判断... |

      输出:

      ```json
      {
          "code": 0,
          "msg": "",
          "result": {
          "author": "",
          "song_name": "",
          "song_url": "",
          "song_silk_base64": "只有这个"
          }
      }
      ```

      示例: `GET http://localhost:3317/silk?url=https://fake.mp3`

