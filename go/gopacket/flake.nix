{
  description = "Network sniffer in Go";
  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-23.05";
  outputs = { self, nixpkgs, ... }@inputs:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
      lib = pkgs.lib;
    in
    {
      formatter.${system} = pkgs.nixpkgs-fmt;
      packages.${system} = {
        default = pkgs.buildGoModule {
          pname = "network";
          version = "0.1";
          src = ./.;
          vendorSha256 = "sha256-sRFJxw6pvcE6KKlDmVvq5QpPkErIdeVh5CjKGUXULAg=";
          CGO_ENABLED = 1;
          preBuild = ''
            GOOS="windows";
          '';
          subPackages = [ "cmd/filter" ];
          buildInputs = with pkgs; [ libpcap ];
        };
      };
    };
}
