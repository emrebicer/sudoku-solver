{
  description = "sudoku-solver go flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.11";
    nixpkgs-unstable.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs = { self, nixpkgs, nixpkgs-unstable, ... }@inputs: 

    let 
      system = "x86_64-linux";
      pkgs = nixpkgs-unstable.legacyPackages.${system};
    in {
        devShells.${system}.default = pkgs.mkShellNoCC {
          buildInputs = with pkgs; [
            (pkgs.python3.withPackages (python-pkgs: with python-pkgs; [
            psutil
            ]))
            go
            gopls
          ];

        };
    };
}
