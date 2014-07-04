with import <nixpkgs> {};
goPackages.buildGoPackage rec {
  version = import ./version.nix;
  name = "direnv-${version}";
  goPackagePath = "github.com/zimbatm/direnv";

  meta = {
    homepage = "http://direnv.net";
    description = "path-dependent environments";
    maintainers = [
      stdenv.lib.maintainers.zimbatm
    ];
    license = stdenv.lib.licenses.mit;
    platforms = go.meta.platforms;
  };
}

