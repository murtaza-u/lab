/*
  {
  description = "Hello Cowsay";
  inputs = {
  nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-22.11";
  };
  outputs = { self, nixpkgs }:
  let pkgs = nixpkgs.legacyPackages.x86_64-linux;
  in {
      formatter.x86_64-linux = pkgs.nixpkgs-fmt;
      packages.x86_64-linux.hello = pkgs.hello;
      packages.x86_64-linux.cowsay = pkgs.cowsay;
      packages.x86_64-linux.default = self.packages.x86_64-linux.hello;

      devShells.x86_64-linux.default = pkgs.mkShell {
        buildInputs = [
          self.packages.x86_64-linux.hello
          self.packages.x86_64-linux.cowsay
          pkgs.lolcat
        ];
      };
  };
  }
*/

{
  description = "Hello Go";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-22.11";
  };
  outputs = { self, nixpkgs }:
    let
      pkgs = nixpkgs.legacyPackages.x86_64-linux;
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";
      version = builtins.substring 0 8 lastModifiedDate;
    in
    {
      formatter.x86_64-linux = pkgs.nixpkgs-fmt;
      packages.x86_64-linux.hello = pkgs.buildGoModule {
        pname = "hello";
        src = ./.;
        inherit version;
        vendorSha256 = "sha256-pQpattmS9VmO3ZIQUFn66az8GSmB4IvYhTTCFn6SUmo=";
      };
      packages.x86_64-linux.world = pkgs.buildGoModule {
        pname = "world";
        src = ./.;
        inherit version;
        vendorSha256 = "sha256-pQpattmS9VmO3ZIQUFn66az8GSmB4IvYhTTCFn6SUmo=";
      };
      packages.x86_64-linux.default = self.packages.x86_64-linux.hello;
      devShells.x86_64-linux.default = pkgs.mkShell {
        buildInputs = with pkgs; [ go gopls ];
      };
    };
}
