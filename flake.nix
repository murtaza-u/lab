{
  description = "Lab experiments";
  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-23.11";
  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      formatter.${system} = pkgs.nixpkgs-fmt;
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          clang-tools
          gcc
          go
          go-tools
          gopls
          nodejs
          nodePackages.pnpm
          nodePackages.vscode-langservers-extracted
          typescript
          nodePackages.typescript-language-server
          nodePackages.live-server
          tailwindcss-language-server
          prettierd
          python3
          kind
          kubectl
          nodePackages.svelte-check
          nodePackages.svelte-language-server
          ispell
          awscli2
          bun
          air
          mycli
          openssl
          nixd
          nixpkgs-fmt
          entr
          yarn
        ];
      };
    };
}
