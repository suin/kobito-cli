require 'formula'

class KobitoCli < Formula
  homepage 'https://github.com/suin/kobito-cli'
  head 'https://github.com/suin/kobito-cli.git'

  def install
    zsh_completion.install 'zsh-completions/_kobito' => '_kobito'
    bin.install Dir['bin/*']
  end
end
