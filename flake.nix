{
  # useful links to understand this file
  # - https://determinate.systems/posts/nix-run
  # - https://nixos.wiki/wiki/Flakes
  # - https://github.com/numtide/flake-utils
  # - https://github.com/edolstra/flake-compat
  description = "Nix flake for TerMImagenator";

  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    flake-compat.url = "https://flakehub.com/f/edolstra/flake-compat/1.tar.gz";
  };

  outputs = { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.${system}; in
      {
        packages = rec {
          termimagenator = with pkgs; buildGoModule {
            pname = "termimagenator";

            version = "latest";
            src = ./.;

            vendorSha256 = "sha256-YRc77w8y6F4B1U3+sh7R4aPpXAvyshRi6uJ0D0smLts=";

            meta = {
              mainProgram = "tmi";
            };
          };
          default = termimagenator;
        };
      }
    );
}
