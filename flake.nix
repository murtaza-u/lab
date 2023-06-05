{
  description = "Lab experiments";
  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-23.05";
  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      formatter.${system} = pkgs.nixpkgs-fmt;
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          go
          go-tools
          gopls
          nodejs
          nodePackages.vscode-langservers-extracted
          typescript
          nodePackages.typescript-language-server
          nodePackages.live-server
          python3
          kind
          kubectl
        ];
      };
    };
}
