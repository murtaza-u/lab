{
  description = "Go with nix flake";
	inputs = {
		nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
		flake-utils.url = "github:numtide/flake-utils";
		gomod2nix.url = "github:nix-community/gomod2nix";
	};
  outputs = { self, nixpkgs, flake-utils, gomod2nix }:
    (flake-utils.lib.eachDefaultSystem
      (system:
        let
          pkgs = import nixpkgs {
            inherit system;
            overlays = [ gomod2nix.overlays.default ];
          };
        in
        {
          packages.gowithnix = pkgs.buildGoApplication {
            pname = "gowithnix";
            version = "0.1.0";
            src = ./.;
            pwd = ./cmd/gowithnix;
            modules = ./gomod2nix.toml;
          };
          packages.hello = pkgs.buildGoApplication {
            pname = "hello";
            version = "0.1.0";
            src = ./.;
            pwd = ./cmd/hello;
            modules = ./gomod2nix.toml;
          };
          devShells.default = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
              gopls
              gomod2nix.packages.${system}.default
            ];
          };
        })
    );
}
