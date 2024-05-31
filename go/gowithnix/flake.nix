{
  description = "Go with nix flake";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
		gomod2nix.url = "github:nix-community/gomod2nix";
  };
  outputs = { self, nixpkgs, gomod2nix }:
  let
    system = "x86_64-linux";
    pkgs = import nixpkgs {
      inherit system;
      overlays = [ gomod2nix.overlays.default ];
    };
  in
  {
    packages.${system} = {
      default = pkgs.buildGoApplication {
        pname = "gowithnix";
        version = "0.1.0";
        src = ./.;
        modules = ./gomod2nix.toml;
      };
      gowithnix = pkgs.buildGoApplication {
        pname = "gowithnix";
        version = "0.1.0";
        src = ./.;
        modules = ./gomod2nix.toml;
        subPackages = [ "cmd/gowithnix" ];
      };
      hello = pkgs.buildGoApplication {
        pname = "hello";
        version = "0.1.0";
        src = ./.;
        modules = ./gomod2nix.toml;
        subPackages = [ "cmd/hello" ];
      };
    };
    devShells.${system}.default = pkgs.mkShell {
      packages = with pkgs; [
        go
        gopls
        gomod2nix.packages.${system}.default
        figlet
      ];
      shellHook = ''
        echo "Welcome to the Nix development environment"
      '';
      IS_NIX_AWESOME="yes";
    };
  };
}
