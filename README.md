# world-parse
This tool transforms a world exported from World Anvil into a lorebook file compatible with SillyTavern

## Quick start
- Download world-parse, you can download the latest version from the releases page.
- Export your world from world-anvil [direct link](https://toolbox.soullink.docker.worldanvil.com/backup/form)
    - Configuration > Open Tools & Advanced Actions > Export world > Export this world
- Run the tool on the world's folder
    - This path is **inside** of the export folder.
    - ex. With export folder `2023-09-06T18\ 33-export`, run the tool like `./world-parse lorebook 2023-09-06T18\ 33-export/worlds/[your_world_here] where/to/put/lorebook.json`

## Details
- Each article gets converted into exactly one lorebook entry, so extra long entries may get truncated according to your SillyTavern world-info settings.
- Backlinks (like `@Name[Person:123abc]`) are replaced with the link's text.
    ex. `@Name[Person:123abc]` becomes `Name`
- The article's named is used as its only tag.
    - Extra tags can be added manually in the SillyTavern UI
- Commas in names will result in SillyTavern assuming that name is two seperate tags
    - This is often desirable e.g. `Eos, Steward of Time` triggers on both `Eos` and `Steward of Time`