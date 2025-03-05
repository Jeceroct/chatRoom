;汉化：MonKeyDu 
;由 Inno Setup 脚本向导 生成的脚本,有关创建 INNO SETUP 脚本文件的详细信息，请参阅文档！!

#define MyAppName "望子成龙小学聊天室"
#define OutputName "chatRoom-v0.2.3-preview"
#define MyAppVersion "0.2.3"
#define MyAppPublisher "HRX"
#define MyAppURL "https://gitee.com/Hu_BanXian/chatroom"
#define MyAppExeName "chatroom.exe"

[Setup]
;注意：AppId 的值唯一标识此应用程序。请勿在安装程序中对其他应用程序使用相同的 AppId 值。
;（若要生成新的 GUID，请单击“工具”|”在 IDE 中生成 GUID）。
AppId={{85DC5F9D-8131-4C57-BE3E-806E4B791E37}
AppName={#MyAppName}
AppVersion={#MyAppVersion}
;AppVerName={#MyAppName} {#MyAppVersion}
AppPublisher={#MyAppPublisher}
AppPublisherURL={#MyAppURL}
AppSupportURL={#MyAppURL}
AppUpdatesURL={#MyAppURL}
DefaultDirName={autopf}\chatRoom
; "ArchitecturesAllowed=x64compatible" 指定安装程序无法运行
;在 Arm 上的 x64 和 Windows 11 以外的任何东西上.
ArchitecturesAllowed=x64compatible
; "ArchitecturesInstallIn64BitMode=x64compatible" 请求
;在 x64 或 Arm 上的 Windows 11 上以“ 64 位模式”进行安装，,
;这意味着它应该使用本机 64 位 Program Files 目录和
;注册表的 64 位视图.
ArchitecturesInstallIn64BitMode=x64compatible
DefaultGroupName={#MyAppName}
AllowNoIcons=yes
; 取消下列注释行，在非管理员安装模式下运行(仅为当前用户安装.)
;PrivilegesRequired=lowest
PrivilegesRequiredOverridesAllowed=dialog
OutputDir=C:\Users\123\Desktop
OutputBaseFilename={#OutputName}
Compression=lzma
SolidCompression=yes
WizardStyle=modern

[Languages]
Name: "chs"; MessagesFile: "compiler:Default.isl"

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked

[Files]
Source: "C:\Users\123\Code\HVVM\chatroom\gin\{#MyAppExeName}"; DestDir: "{app}"; Flags: ignoreversion
Source: "C:\Users\123\Code\HVVM\chatroom\gin\chatroom_debug.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "C:\Users\123\Code\HVVM\chatroom\README.md"; DestDir: "{app}"; Flags: ignoreversion
Source: "C:\Users\123\Code\HVVM\chatroom\gin\dist\*"; DestDir: "{app}\dist"; Flags: ignoreversion recursesubdirs createallsubdirs
; 注意:  在任何共享系统文件上不要使用 "Flags: ignoreversion"

[Icons]
Name: "{group}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"
Name: "{group}\{cm:UninstallProgram,{#MyAppName}}"; Filename: "{uninstallexe}"
Name: "{autodesktop}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"; Tasks: desktopicon

[Run]
Filename: "{app}\{#MyAppExeName}"; Description: "{cm:LaunchProgram,{#StringChange(MyAppName, '&', '&&')}}"; Flags: nowait postinstall skipifsilent

