{
  description = "Setting up Go TEMPL + Tailwind + HTMX";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-23.11";
    templ = {
      url = "github:a-h/templ";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };
  outputs = { self, nixpkgs, ... }@inputs:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
      templ = system: inputs.templ.packages.${system}.templ;
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
          tailwindcss-language-server
          prettierd
          air
          (templ system)
        ];
      };
    };
}
