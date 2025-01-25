<div id="header">
   <h3 align="center">
      <span style="color:#fa6546">Mangadex</span> 
      <span style="color:#e4dddb">Downloader</span>
      <br>
   </h3>
</div>

>An (unofficial) CLI to download manga chapters from [Mangadex](https://mangadex.org/) with just one command, with the possibility of obtaining them in .zip and .cbz formats.

---

![download-cmd](https://raw.githubusercontent.com/Otarossoni/mangadex-downloader/master/assets/download-cmd.png)

## Usage

The use is done entirely through flags, from the indication of the manga to be downloaded, to the path to generate the final file. The manga can be indicated in two ways, by the `--url` flag, being the complete URL of the manga in Mangadex, or by the `--mangaId` flag, being only the UUID of the manga informed. See the example below:

```bash
./mangadex-downloader --url https://mangadex.org/title/6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b/tokyo-ghoul
# or
./mangadex-downloader --mangaId 6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b
```

All flags have aliases, see the example above only using them:

```bash
./mangadex-downloader --u https://mangadex.org/title/6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b/tokyo-ghoul
# or
./mangadex-downloader --m 6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b
```

You can review all flags through the command `./mangadex-downloader -h`:

![help-cmd](https://raw.githubusercontent.com/Otarossoni/mangadex-downloader/master/assets/help-cmd.png)

#### Flags

All parameter options that the CLI can receive:

- `--url`: Specifies which manga to download (required if the "--mangaId" flag is not provided).
- `--mangaId`: Specifies which manga to download (required if the "--url" flag is not provided).
- `--chapters`: (Optional) Specifies the range of chapters to be downloaded.
  - Ranged by "-", segmented by ";"
- `--language`: (Optional) Specifies the languages of the chapters.
  - en(default), pt-br, es-la, pl, cs, uk, it, vi, hu and others accepted by Mangadex
- `--extension`: (Optional) Specifies the extension of the final compressed file.
  - .zip(default), .cbz
- `--outPath`: (Optional) Specifies the path to generate chapter files.
  - Same path as executable (default)
