require "language/go"

class Sapin < Formula
  desc "Draw a beautiful christmas tree in ascii"
  homepage "https://github.com/moul/sapin"
  url "https://github.com/moul/sapin/archive/v1.1.0.tar.gz"
  sha256 "b46f4eea39b383825d2646bdbb63f1e98775027de83625f1e9362eefbf112161"

  head "https://github.com/moul/sapin.git"

  depends_on "go" => :build
  depends_on "godep" => :build
  depends_on "jq" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["CGO_ENABLED"] = "0"
    ENV["GO15VENDOREXPERIMENT"] = "1"
    ENV.prepend_create_path "PATH", buildpath/"bin"

    mkdir_p buildpath/"src/github.com/moul"
    ln_s buildpath, buildpath/"src/github.com/moul/sapin"
    Language::Go.stage_deps resources, buildpath/"src"

    system "godep", "restore"
    system "make", "build"
    bin.install "sapin"

    # bash_completion.install "contrib/completion/bash/sapin"
    # zsh_completion.install "contrib/completion/zsh/_sapin"
  end

  test do
    output = shell_output(bin/"sapin --size=1")
    assert output.include? "
   *
  ***
 *****
*******
   |"[1..-1]
  end
end
