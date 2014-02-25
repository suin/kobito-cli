require 'formula'

class KobitoCli < Formula
  homepage 'https://github.com/suin/kobito-cli'
  head 'https://github.com/suin/kobito-cli.git'

  def install
    bin.install Dir['bin/*']
  end
end
