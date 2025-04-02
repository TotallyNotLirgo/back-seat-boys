{
    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    };

    outputs = { nixpkgs, ... }:
        let
            system = "x86_64-linux";
            pkgs = nixpkgs.legacyPackages.${system};
        in
            {
            devShells.${system}.default = pkgs.mkShell {
                packages = with pkgs; [
                    go
                    tparse
                    gopls
                    air
                ];
                shellHook = ''
                source ~/.bashrc
                cd ..
                tput setaf 2
                echo "Happy hacking!"
                tput sgr0
                '';
            };
        };
}
