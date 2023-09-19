%define debug_package %{nil}
%global command_name suitectl

Name: suitekit
Version: %{version}
Release: 1
Summary: SuiteKit rpm pkg
Source0: go.tgz
Source1: source.tgz

License: Proprietary

Requires: bash-completion

%description
SuiteKit is a tool written in Golang, for init and deploy ITSMA services, and CDF as well.

%prep

%setup -q -c -a 0  # This will extract go1.20.7.linux-amd64.tar.gz
%setup -q -c -a 1  # This will extract suitekit.tgz


%build
# Add the path of Go to the PATH
export PATH=$PATH:%{_builddir}/%{name}-%{version}/go/bin
BUILD_DATE=$(TZ=Asia/Shanghai date +"%Y-%m-%d_%H:%M:%S_%Z")
CGO_ENABLED=0 GOOS=linux go build -a -gcflags="all=-N -l" -ldflags "-extldflags '-static' -X suite-kit-go/cmd.buildDate=$BUILD_DATE" -o %{command_name} main.go

%install
mkdir -p %{buildroot}/%{_bindir}
install -m 755 %{command_name} %{buildroot}/%{_bindir}/%{command_name}

%post
sed -i '/%{command_name} completion bash/d' /etc/bashrc
echo "source <(%{command_name} completion bash)" >> /etc/bashrc

%files
%{_bindir}/%{command_name}

%changelog 
* Wed Aug 13 2023 Haoyun Ji <hji2@opentext.com> - 0.0.5
- Add debug info.
- Add command example to silent-install
* Wed Aug 13 2023 Haoyun Ji <hji2@opentext.com> - 0.0.4
- Optimize executeSilentInstall() function
* Wed Aug 13 2023 Haoyun Ji <hji2@opentext.com> - 0.0.3
- Fix rpm actions.
* Wed Aug 13 2023 Haoyun Ji <hji2@opentext.com> - 0.0.2
- Update go version to 1.21
* Wed Aug 13 2023 Haoyun Ji <hji2@opentext.com> - 0.0.1
- Support silent-install