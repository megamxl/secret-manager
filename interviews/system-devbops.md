

   0:03
Nachdenken.

   0:03
Ha ha ha.
Ah.

   0:06
So, was ist deine aktuelle Position?

   0:10
Also, ich bin Linux System Engineer.
Und hauptverantwortlich für LDAP *** Automation Platform, genau.

   0:20
Ja, und wie lang machst du das schon? Schrägstrich bist generell in der IT unterwegs.

  0:25
Mittlerweile.
Gute 10 Jahre schon.

   0:30
Wunderbar.
Ich glaub, bist du jeden Tag mit Secret Management schrägstrich Application Secret Management umgeben oder ist das eher das, was du selten machst?

  0:42
Da bin ich eigentlich fast jeden Tag damit betroffen.

    0:46
Verstehe, bist du mit dem aktuellen Approach zufrieden? 

   0:51
Ähm, mit den geplanten Approach, äh, ***.

    0:58
Verstehe, hast du jemals den External Secret Operator verwendet bis jetzt? Das ist ein, ja, das ist ein komponiertes Ding.

   1:04
An external secret operator, oder wie? ***.

   1:09
Sonst wunderbar, wunderbar, dann wirst du ihn gleich kennenlernen. Der ist jetzt das ganze Konzept wichtig. Genau, Hallo ***, was ist deine aktuelle, was ist deine aktuelle Position?

  1:09
No, then can you know then?
Mhm.

   1:17
Hallo Maximilian.

  1:19
Mhm.

   1:21
Ich hab grad nachgeschaut in meiner Signatur. Das ist sehr lustig. Ein Devox Engineer ist die Bezeichnung bei uns.

   1:26
OK.
 Auch seit 10 Jahren wunderbar gut. Wie zufrieden bist du mit dem Secret Management, was du zur Zeit anwenden darfst? ***

   1:37
Nee, was? Das wird dich zum Lachen bringen, aber ich bin nicht so unzufrieden. Natürlich geht es immer besser. Du bist ja ein Grund, also die Wahrheit ist, du bist ja da, um das auch zu verbessern. Das heißt, wir beschreiten eh gute Wege, auch wenn wir sie langsam gehen.

   1:56
Versteh.

   1:56
Nicht auf einer Skala von 1 bis 5 würde ich um 3.

   2:01
Hast du schon mal was im External Secret Operator gemacht?

   2:04
Ah.

   2:05
Passt gut. Also, was habe ich jetzt in meiner Masterarbeit gemacht? Also, es gibt in Kubanetis ein paar Wege, wie man Dynamic Secret Management sozusagen gut applian kann. Und dafür müssen wir jetzt mal verstehen, was ist das sind Secrets, Dynamic Secret Management U.S.W.
Für Applications hat man Secrets sehr gerne in Vaults heutzutage und das ist eigentlich der vorgegebene Weg, weil das ist ein abgesichertes System. Das ist ziemlich hardened, da kannst du halt den Lifecycle von Secrets managen, sozusagen. Aber dieses Delivery dann von Secret zur Applikation ist sozusagen noch immer eine relativ offene Gap, wo es einige Tools gibt.
Die sozusagen sich zwischen Linux und Kubanetisch stark unterscheiden und das ist halt einfach. Es fügt das führt zu Secret Sprawl und generell Secret Management per se auf Dynamic Ebene auf Linux Hosts ist jetzt nicht wirklich etwas, was.
sonderlich gut, wenn Diagnostik umgesetzt wurde. Das heißt, ein Secret Store per se ist halt einfach ein Ort dann mit einer REST A. P. I., wo ich ein Secret ablegen kann und eine gewisse Lifecycle Management habe. Was dann halt wieder Applikationen abholen können. Genau. Dexter Secret Operator ist jetzt sozusagen unsere Bridge, unser Proxy, der redet mit
Zum Beispiel Asia Keyboard.
Also, dem konfiguriere ich eine Secret Store Connection. Also, wie authentifiziere ich mich beim Vault und wie soll das Secret, was am Ende des Tages da ist, aussehen ohne Werte reinzuschreiben? Und der verbindet sich mit dem Vault. Kreiere Template, also holt sich die Werte. Templated the Secret legt's ab in Kubernetes.

  3:21
Mhm.

   3:35
Und dann kann seine Applikation mounten, also sozusagen generiert ein File, was irgendwo rumliegt im ICCD, was sozusagen dann als als Secret Objekt in Kubernetes verfügbar ist.
Das ist eigentlich ein supercooler Approach, weil er ermöglicht es dir, runtime undependent zu sein. Das heißt, das passé, passé, wenn Asia morgen nicht geht und das Credential für deine interne Datenbank noch mehr funktioniert, startet deine Applikation immer noch. Was für viele Applikationen wichtig sind, nicht jede Applikation kann sagen, nur wenn die Asia da ist, kann ich starten.
Genau, und es ermöglicht dir auch sozusagen ein Vendiagnostic-Approach. Das heißt, du kannst heute affault sein, morgen auf einem anderen und übermorgen auf einem ganz anderen System und die Migrationspfade sind relativ leicht. Das heißt, du bist nicht gelockt auf den Secret Store, den du jetzt mal definierst.
Genau, und aber es gibt irgendwie für Linux nichts, was konzeptionell so einfach funktioniert. Auf Linux gibt es eher so Sachen wie "Ich löse das mit Ensible, ich habe ein eigenes Skript, was mir das macht" oder "Ich habe irgendeinen Chronchop, der irgendwas macht und der fällt silent herum", und das ist sozusagen alles nicht wirklich.
ganz sinnvoll gelöst. Demnach habe ich für meine Masterarbeit einen hybriden Secret Management Workflow definiert. Der besteht daraus, dass wir den externen Secret Operator auf Kubernetes haben und ich habe so einen kleinen Linux Agent gebaut, der das Gleiche halt auf Linux abbilden soll.
als Prototyp. Genau, und der Workflow ist, ich habe halt ein Git Repository, in das kann ich halt meine Secret Configurations und meine Store Configurations einchecken. Deine CISD Pipeline nimmt die dann, schmeißt die dann auf entweder den Kubernetes Server oder auf meinen Linux Host drüber und der und in bei den Fällen kümmert sich der Operator darum, diese getemplateten
Comfix in Files aufzulösen.
Thank commerce of that to claw.
Genau, wie würde sozusagen dieser Agent ein bisschen genauer ausschauen? Cognitis gibt dir ja ein paar coole Eigenschaften, wie eine erreichbare API, Authentication, Authorization und so weiter und halt den Datastore mit ETCD. Das heißt, dieses diese Spinary muss sozusagen einiges.
Einige der Features abbilden halt nicht in der Qualität wie Open, also wie in Kubernetes, aber du brauchst halt gewisse Sachen. Demnach exposed er eine REST API und der hat sozusagen 2 REST API Pfade, einmal für Stores, also für die Connection zu Secret Stores und eine für
die Konfiguration von den Secrets. Das legte sich in einen SQLite Database File ab, was er am Host generiert. Das heißt, da stehen auch die Login Crudentials zum also zum Secret Store drinnen. Genau, das SQLite Binary ist gepackaged im Go Binary, das heißt, es gibt keine Dependency, dass du das davor installieren musst.
Genau, und was er auch noch kann, er schreibt jetzt seine Tomicly. Was ich damit genau meine, komme ich später noch dazu. Das ist ein kleiner Workflow, den ich mir überlegt habe. Und es gibt einen Slash Metrics Endpoint, um halt sozusagen, nachdem das ein massiv distributed System ist, die ganzen Uptime Alerting und Sachen umsetzen zu können.
Noch dazu, wenn ein Secret nicht rollierbar ist, wird es auch in diesen Endpoint reingeschrieben. Und darüber kann man sich dann Alerts generieren und sozusagen sicherstellen, dass wenn etwas fehlschlägt, dass man es relativ schnell erwischt und nicht erst so ein paar Wochen drauf kommt, dass dieses eine File eigentlich nicht richtig rotiert wird.
Genau sozusagen, so viel einmal zum Solution Approach und es gibt dann halt eine C.I.C.D. Pipeline oder einen DevOps Engineer, je nachdem, wie man es halt aufsetzen möchte, der sich halt dann um das Management von diesen Konfigurationen kümmert, von den Secrets der Connections sowie von den C.I.C.D. oder halt.
Sehr ICT mäßig genau.
Ich glaub, das sollte auch noch klar sein. Demnach gibt es jetzt ein 2 große Konfig Type, einen Secret Store. Das heißt, in dem steht drinnen, wo ist der Vault, wie authenticate ich mich mit dem Vault und beim Anlegen, schau da, ob er sich mit den Potentials überhaupt anmelden kann, weil sonst hast du halt eine Connection, die vielleicht gar nicht funktioniert und hast dann später ein Problem.
Genau, und dann gibt es Managed Files, was in Kubernetes the Secret ist, ist jetzt sozusagen eine Dateipfad, den ich angeben kann. Dann habe ich einen Templating, so wie beim Excel SECret Operator, mit dem kann ich sagen, an der Stelle replace diesen Wert bitte mit dem Wert aus dem Vault.
Ich habe eine Time to Lift, das heißt, der sagt mir so und so lang darf dieses File existieren und darüber hinaus ersetzen wir es, so lang es funktioniert. Und wenn es ein Jason Jamel AN vor der Pempal ist, schau da an, ob das File per se valide ist, weil rein tödlich könntest du ja auch im Secret Store.
Wenn du in Jason hast, in deinem Passwort vielleicht einen Cotation Mark drin haben und dann hättest du ein unvalides File und dein System wird nicht mehr hochfahren. Das gleiche Problem hast du beim External Secret Operator, das hab ich nur einfach abgefangen in meinem Workflow, dem nachher wenig passiert.
Genau, noch mal kurz zusammengefasst, es gibt eine CRAD API für Secret und Store Management sowie in Kubernetes Basically. In mit GET kriege ich sozusagen die Konfigurationen für meine Secrets zurück und meiner Stores, aber natürlich sind die auf Indication Daten redacted, das heißt.
Die werden nicht mehr an den User zurückgegeben. Generell werden keine Secret Contents, also die Files am Host sind nicht auslösbar über die API. Die existieren nur im Host. Es gibt den Audit Lock, dann gibt's die Time to Lift, den Matrix Endpoint und man muss explizit in der Comfig erlauben, wo er hinschreiben darf.
Weil, wenn du leicht rennst, vielleicht läuft es ja mit einem User, der mehr privileged ist als die Folder, die er hat. Demnach gibt's da noch mal 'n Sicherheitsmechanismus, dass ich im Agent direkt konfigurieren muss, welche Pfade er als als Root hat, also schreiben darf dort.
Genau nachdem der Agent per se authentication mäßig so gelöst sein muss, dass auch im Internet hängen könnte, weil das erlaubt dir Kubernetes auch, habe ich mich hier für einen Layer-Approach entschieden, dass es MTLS gibt, der als erstes halt checkt, ob der Client überhaupt authorized ist, mit dir zu reden. Und dann kannst du noch einen JVT mitgeben.
Und dieser JVT wird dann auch noch validiert über über die ORF Protokolle, also mit einem Public Key, den er sich per Rest holen kann von irgendwo, was sozusagen dafür sorgt, dass du wenig Exposure hast, dass sozusagen nur dass du die TLS Settings musst du irgendwie auf den Host bringen.
Und ein Workaround, wie man das machen kann, komme ich dann später noch dazu, aber da musst du halt die Initial Root Codentials für den Agent Pro Host rüberkriegen. Aber sobald du die dort hast, kannst du eigentlich MTLS und JVT betreiben, ohne großartige Probleme. Genau, kommen wir zum Atomic Fileswap. Der Atomic Fileswap ist einfach ein kleiner Prozess, wo wir halt checken, ob
Workflow-Algorithmus, je nach dem checken. Darf ich da reinschreiben? Ist der Content, den wir gebaut haben, ein File, das ich vom Format kenne? Da probieren wir es zu marsheln. Wenn das funktioniert, kreieren wir ein Tempfile. Dieses Tempfile wird dann mit den richtigen Permissions gesetzt. Dann schließen wir dieses Tempfile und dann schauen, dass das existiert.
Und dann machen wir ein OS. Punkt. Rename. Das ist der einzige Atomic Process in Linux, um Files zu transferieren, wenn mich nicht alles täuscht. Und könnte dich sicherstellen, dass ein File, wenn es geswappt ist, hundertprozentig geswappt ist und keine Zwischenstände oder Unfinished Files irgendwo existieren, weil DiscSpace ausgegangen ist oder.
und so weiter. Genau.
Wie hätte ich mir jetzt vorgestellt, äh den Excellent Secret Operator kann ich auf Kubernetes relativ einfach hauen, das sage ich einfach. Cube C. D. L. Apply oder O. C. Apply und geben halt das Manifest. Das Gleiche hast du halt nativ auf Linux nicht. Könntest natürlich probieren, irgendwie so was wie Talos Linux zu erweitern, dass du sozusagen eine Rest IP für dein Linux hast, aber.
Das wird ja auch nicht sehr viel Spaß machen. Demnach habe ich mich für Ensible entschieden. Genau, weil Ensible relativ ähnlich ist zu CubeCDL, muss ich halt mit Linux bisschen auseinandersetzen. Genau, und das macht der Installer. Er kreiert den System User, eine Secret Gruppe, die Pfade, die er braucht.
Er templated dir die Konfiguration, macht das System D-Service und lädt das Binary am Host runter.
Genau.
Jetzt kurz zur Konfiguration, was kann ich dem Agent jetzt mit konfigurieren? Ich kann ihm sagen, wohin darfst du schreiben, auf welchem Port bist du erreichbar, wie schaut denn Logging aus, wo legst du bitte genau das Audit Logfile hin, wo liegt die SQLite Database. Dann kann ich im TLS mäßig halt auf und abdrehen, brauche ich halt das Certain das Keyfile.
das Client CA File, wenn ich MTLS haben möchte und halt und halt für Security können wir uns zwischen Non MTLS und JVT empfehlen. Das sicherste ist natürlich, wenn wir das Client CA File angeben und für die Authentication dann auch noch JVT verwenden. Rein kritisch könnte es doch nur MTLS haben und sagen, solange der
Request M.T.L.S. mäßig signiert wurde, passt das? Das kommt halt drauf an, welche Thread Exposure du hast und welches Thread Level du dem Ganzen eingibst. Genau, und bei JVT kann ich entweder dieses Shared Secret machen, die H. MAC Variante, wie wir es kennen, von JVT, oder ich kann eben die JVCase URL angeben, also von einem Identity Provider wie KeyClog.
wenn du einen Raym hast, hast du sozusagen Keys, die sozusagen wo die JVTs, die von KeepClock kommen oder von einem anderen OVP-Provider halt mit deren private Keys signiert wurden und du hast halt die Public Keys, um zu validieren, was sozusagen besser in dem Fall ist, weil du brauchst halt dich nicht um das Secret der JVK, also die JVTs kümmern.
sondern die holt sich der Agent dann, wenn er sie braucht. Hast du halt weniger, was du halt irgendwo encrypt musst und irgendwo operieren musst. Genau, kommen wir jetzt zum Configuration Example. Hier sehen wir es, wie es der XNL Secret Operator gelöst hat. Das hier jetzt von Asia Keyboard, das sagt hier die Tenant ID, hier die Vault URL und dann referenziert er die Secrets.
nachdem du in Kubernetes per se halt noch mal diese Secret A. P. I. hast, kannst du die hier referenzieren, weil das heißt, du erstellst sozusagen als erstes ein Secret mit den Credentials zum Store und hast dann sozusagen eine Store Definition, die du sozusagen ins Git einchecken kannst.
Das geht bei meinem Prototypen noch nicht, der geht relativ gleich. Aber anstatt, dass du halt hier diese Werte als irgendwo anders hast, hast du die halt hier eben auf Block. Das heißt, die müssten wahrscheinlich in einem Extra Repository oder halt von einem oder halt händisch irgendwo gewartet werden.
Für den Prototypen, weil genau. Dann kommen wir auch schon zum Configuration Example und hier sehen wir sozusagen, wie das dann funktioniert. Im External Sequel Operator hast du halt die Definition von wo, also wie du dich connecten möchtest. Dann hast du halt hier dein Kontext. Also dann hast du halt hier das, was du templaten willst, in unserem Fall ein ganz einfaches Kontext.
Zeil mit User Password und URL. Das heißt, wir sehen hier .db Password haben wir hier unten.db Passwort.
Jetzt haben wir Punkt dpuser sorry, siehe wird unser DB Username. Etwas relativ ähnliches habe ich da drüben in meiner Konfiguration nachgebaut. Du musst also auf Linux musst du halt angeben, wo du hinschreiben willst, die Datei. Dann schau, dann hast du halt auch die gleiche Rero-Time, bei denen heißt Refresh Interval.
Dann den Storename, mit welchem Store du connecten möchtest, hier den Template und dann statt diesen Bulky Teil habe ich mich hier für ein bisschen ein einfaches System entschieden. Du hast dieser Wert ist die also du hast den Mapping von hier auf hier zu hier, das heißt.
Das Das löst du dir auf, sozusagen intern. Dann, das ist dein Templating Key und das ist der Templating Value und dieses holt er sich aus dem Store. Das heißt, wenn du jetzt mit dem Store arbeiten würdest, der hierarchische Pfad unterstützt wie der Haschikop Keyboard, dann hättest du hier dev
TB Username, währenddessen du halt beim Asia Keyword keine Error hier abbilden kannst. Demnach ist das nur ein TB Username. In dem Fall genau.
Fragen wir mal zu dem Ganzen.
Irgendwelche Anmerkungen? Irgendwas schon einmal oder seid ihr einfach nur froh, dass ihr da sitzt?

  15:04
Na ja, also.

   15:04
Alles nicht.

  15:06
Da ist ganz schlüssig und.

   15:09
Wunderbar.

  15:09
Eigentlich ja, eine gute Idee.

   15:12
Wunderbar, dann hab ich ein Survey für euch.
D d, was ich euch bitten würde zum Ausfüllen und dann habe ich noch ein paar offene Fragen und dann sage ich auch schon vielen Dank für eure Zeit. Bitte fragt mich, Fragen zu fragen, wenn ihr es nicht versteht. Das ist der Grund, warum ich das sozusagen in dem Interview in dem Part mach, weil sozusagen die Fragen vielleicht entwickelt ist.
Formuliert sind und man sie nicht hundertprozentig versteht. Genau, und noch zu dem Thema zum Agent, was mir aufgefallen ist, bevor er das beantwortet. Was man rein typisch machen könnte mit seinem Agent, wäre das Cert-File, das Key-File und das Client CA-File in der Bootstrap-Phase.

  15:33
Mhm.

   15:50
Sozusagen mal mit einem Credential befüllen, was eine Time to Live von 3 Tagen hat also halt ein Certificate, was halt nicht lang lebt. Und dann könnte der Agent 33 selber diese Files überschreiben und sich sozusagen selber Secure halten über das.
Verfahren, nur dass ihr das sozusagen wisst, als Input, was man mit dem Agent machen könnte. Also, man kann halt jedes File überschreiben und nachdem das halt ein File wie jedes andere ist, insofern der User berechtigt ist, könnte er sich sozusagen selber rollieren.

   16:20
Aha.

   16:21
Genau.
Dann wünsche ich viel Spaß beim Form.

   16:38
Der zweiten Frage fragst du uns nach einer Doku, die gibt es, oder?

   16:43
Ja, die Doku, die gibt es per se. Schwupp, ich vergessen reinzuschicken. Das ist eben Git Repository.

   16:55
Ich vertraue der Ehe.

   16:59
Ich hab das Repositor eh schon mal geschickt, *** also.

  17:02
Gelesen haben wir es natìrlich auch, die Dokumentation, ja.

   17:04
No, no.

   17:05
Genau, also Max, mach dir keine Sorgen, ich hab mich echt ausführlich damit beschäftigt.

   17:07
But.
Bei euch passt alles. Ihr habt keine Fragen?

  19:28
Der wär nett, äh, äh, doch jetzt steht ISO.

   19:32
External secret operator.

  19:34
Ah, OK.

   19:39
Ich hab grad Elder Scrolls, da muss ich meine Antwort nochmal.

  19:42
Elder Schools online.

   20:08
Das ist gemein, wenn du compare to existing approaches on führst bei uns du.

   20:13
Na ja, irgendwie muss man ja seine seine Solution verkaufen.

   20:20
Huh.

   20:24
Hier ist es so wie bei den ganzen Eissachen. Ich verkaufe nur die Schaufeln. Ich könnte selber dann Gold schaufeln gehen.

   20:54
Versuche ein bisschen ernst zu bleiben bei den Frauen.
Notation Shadows mit die organizational Constraints du bist echt gut.
Ich seh dich nicht.

   21:05
Das Interview ist zu 50% für die Masterwelt zu 50%, um den Composure zu testen, ***.

   21:10
Ha ha ha.

  21:44
Na, ein paar Sachen schlechter.

   21:48
Ja, wie? Was brauchst du, ***? Ich suche auch schon.

   21:52
Du, mir ist es komplett egal. Das ist eine wissenschaftliche Arbeit. Es muss nicht rauskommen, dass das Beste ist. Aber wenn du etwas Unzufriedenes ist, ist es besser, wenn das acknowledged wird und man dann in der Future Work darauf hinweisen kann, dass man das noch verbessern könnte.

  21:58
Mhm.

   22:04
Aber bevor du die Transkription eingeschalten hast und uns gesagt hast, wir sollen alles auf Super tun, hat das anders gelaufen.

   22:13
Das kommt raus. Das streich ich raus.

   22:13
OK.
ha ha ha

   22:18
Nein, Spaß.

  22:19
Also, ich bin durch.

   22:20
Ja.

   22:21
Wunderbar, gut. Dann kommen wir jetzt zu ein paar offenen Fragen, die mich interessieren. Ich werde euch nicht alle stellen, die ihr da seht, sondern weil es interessiert mich halt aufgrund eurer Erfahrung gewisse Fragen.
Genau, was seht ihr sozusagen als Vorteil von dem Agent?

  22:40
Also, wenn ich anfangen soll.
Wir haben im Linux-Umfeld ja immer das Problem, dass wir keine Ahnung, irgendwelche Skripte oder was auch immer irgendwo laufen haben, wo es irgendwie ein Secret brauchst oder ein Passwort brauchst und nachher liegt es oft dann irgendwo in einem File oder Verstecken.
File oder nur als als Hash, ähm, Wert irgendwo herum, was natìrlich jetzt als secret Gesicht net so toll ist und unser Approach dahingehend ist halt alles was ähm
Skripte und KoSan, die Passwörter brauchen zentral zu managen über die *** Automation Plattform und mit dieser haben wir Integration zu dem *** Server, den wir halt in der Firma verwenden, implementiert. Sprich, wir fragen die.
Die *** dann dynamisch aus dem *** Server ab und und dann die dann injecten in die jeweiligen Skripte.
Der Vorteil von dem Agent sehe ich da natürlich, dass das unabhängig von einer zentralen Plattform laufen würde, zumindest jetzt von einer zentralen Jobplattform, weil sie dann der lokale Agent am Server drum kümmert.

   23:54
Wunderbar, danke.

   23:59
Genau dazu noch, das ermöglicht uns überhaupt, diese hybriden Umgebungen zu betreiben, ordentlich so und dann gemeinsamen Standard quasi aufrechtzuerhalten beziehungsweise in unserer World, wenn wir benutzen zwar den *** Server,
Manuell möchte ich sagen, aber dass man auch mal automatisiert das in irgendeiner Form dann oder die Applikationen selber das benutzen können. Bei uns hat das ja bis jetzt nicht mit dem *** Server. Das heißt, es ist überhaupt ein Ablement.

   24:30
Verstehe, wunderbar. Wie seht ihr den Tradeoff, dass das Secret als Plaintextfile herumliegt, als Acceptable oder ist das für euch ein Roadblock?

   24:43
Also, Entschuldige, bist du bereit? Bei uns ist sie, du kennst die Diskussionen eh, Max. Bei uns ist es so, durch das, dass wir heute im *** Bereich quasi uns ein bisschen auf die Firewalls verlassen und diesen abgeschotteten Bereich in verschiedenen Formen.

   24:45
Ja, du beginn ruhig.

  24:47
Mhm.

   24:50
Mhm.

   25:01
Heute ist für ein also für meinen Bereich für ein akzeptables Risiko oder ein Trader of der Halt.
Ja, der, der halt hinnehmbar ist, glaube das kann man ausführlicher diskutieren, aber in Wirklichkeit haben wir andere Mechanismus Mechanismen, die dem noch vorgreifen und und insofern will ich es akzeptieren und sagen, es ist gut.

  25:26
Ja, ich sehe das schon eher kritisch, aber.
Das ist dann die Frage, wie wie lang das Pfeil dann dort liegen muss, wo die Sigrid dort steht, wenn das dann wirklich nur für für der Zeit, wo die Sigrid gebraucht wird, dann abgelegt wird und dann wieder entfernt wird, quasi nur zur gewissen Laufzeit.
Dann würde das O. K. finden, wenn das File dann immer dort liegt. Sehe das schon eher als kritisch, weil wir oft so Angriffssektoren, wo dann User übernommen werden und nachher passt bei den File Permissions was nicht korrekt.
Und dann kann ein User das Sicket File lesen, das eigentlich nicht lesen sollen könnte. Ich weiß, du hast das eher mit den Alarmierungen und und und Sachen halt eingebaut, dass das auch Secure ist, also zumindest bei dem, was du vorher gesagt hast.

   26:09
Ja.

  26:18
Aber ja, das ist das, was da da da genau aufpassen.

   26:23
Verstehe, nein, die Files liegen sozusagen permanent herum, aber sie werden rolliert. Das heißt, das heißt per se, die um den Attack Vector dieser Files zu ändern, hast du halt normalerweise dann eigentlich ein Secret, was nur 12 Stunden lebt und hast halt alle 12 Stunden. Aber natürlich, sobald irgendwas mit dem Host ist und irgendein Host.

  26:26
Yeah.
Ja.

   26:41
eine Host Übernahme ist, sei es Root, sei es irgendwas anderes, sei es du kriegst die Disc von dem System irgendwie gedumpft, hast du sozusagen per se verloren. Aber genau deswegen hast du halt auch diesen diesen Vault per se, dass du halt den Zugang dazu revoten kannst und alle Sequels rollieren kannst.

  26:43
E.
Ah.

   26:58
Und da hilft dir halt auch wieder das, dass das im Kit storen kannst und rein deutlich auf einen neuen Host haben kannst, alles hinschießt und du hast gleich wieder den. Also, du kannst deinen ganzen State in Kit abbilden, was das angeht. Aber das ist definitiv genau darum, frage ich das, um halt damit.

  26:58
Going on there.

   27:13
Noch Sachen für Future.

   27:14
Verbesserung zu unserer Linux Welt ist es jetzt schon, wenn ich mir anschaue, wie unsere *** oder die Secrets gemanagt sind in verschiedenen Files, dann ist ist das Problem, wenn wirklich Maschine übernommen ist. *** Also dann haben wir andere Probleme in Wahrheit, wie das.

  27:15
Ja.
Ja, ja, genau. Also, das, dass das eine Verbesserung ist, ist ist ja logisch. Und das und ich glaub die die Rollierung, wenn du sagst, du rollierst das einmal am Tag oder so, dann Mitte geht das das Aids sehr stark.
Weil dann hast du nicht so eine lange Exposure.

   27:46
Ja, genau.
Ja, also natürlich, wenn jetzt ein, wenn jetzt ein Intruder unerkannt im System ist und jeden Tag dann das neue File ausliest, dann das ist ein anderes Problem, aber.

  27:51
Genau.
Ja, das ist, das ist, das ist eh logisch. Also, du kannst das, wenn er mal so weit ist, das am System ist, ist eh, aber man muss ihm halt immer wieder diese, man muss ihm das Leben so schwer wie möglich machen, weil irgend irgendwelche Möglichkeiten gibt es immer, aber

   28:02
Just.
Yeah.
Da genau.

  28:12
Je schwerer was es hat, desto besser.

   28:12
The.
Secret Management war jetzt natürlich die erste Idee von diesem Workflow und diesem Agent. Könnt ihr euch vorstellen, dass dieser Agent auch noch in anderen Bereichen im Linux-Spektrum einsetzbar wäre für gewisse Tasks?

  28:33
Zur Frage.

   28:35
Mit mit wenig Umbau oder Erweiterungsaufwand oder wie meinst du da?

   28:39
Das ist per se egal, es geht da einfach um Ideen, die man vielleicht in Zukunft exploren kann. Eine, die öfters genannt wird, wäre halt generell nicht nur Secrets, sondern auch genau Konfigurationen in Git Templateable machen, der das dann sozusagen pro Host auflöst.

   28:48
Konfigurationen.

   28:57
Muss man dann immer wissen, wie dynamisch braucht man das und ob da nicht mit J2 Ninja und dem Template Ding reicht.

   29:03
Ich wollt Grad sagen, wenn du die einzige schon in Benutzung hast, dann.

   29:07
Ja, kommt halt, wie gesagt, drauf an, wie sehr du also mit welchem Approach du fahren möchtest und wie sehr die Configuration Changes einen neuen Ensible Road triggern sollen. Aber ja, genau, wenn euch nichts einfällt, ist es auch O. K., das fahr ich immer, weil
Können hier coole Ideen rauskommen?

  29:24
Ja, aber das, das wäre, das wäre mir einfallen mit Konfigurationssachen, Rolladen, wenn es jetzt nicht so eine große oder das eher im Applikationskontext siehst und sagst, du hast deine Konfig, das die sich oft ändert, dass du dann auch in den Tool mit integrierst, dass da zusätzlich dann noch ein Ensible Workflow auch.

   29:36
Uh.

   29:42
Mhm.

  29:42
dann nun dazu brauchst.

   29:44
Mhm. Was mir auch noch über andere und durch mein eigenes Nachdenken eingefallen hier wäre halt sozusagen den halt wirklich stark für die Rollierung von von von von Certificates zu verwenden. Das ist sozusagen deine Webserver und so weiter eigentlich sozusagen sehr auf den neuesten Stand heben kannst und T. L.
S mäßig, indem du einmal in der Woche oder einmal am Tag das TLS Secret endest.
Weil das kann er ja auch anstatt nur Application kann das sozusagen auch Richtung Infrastructure Secret Management verwendet werden.

   30:10
Yeah.

  30:14
It.
Das ist eine sehr gute Idee. Ja, also ich meine, viele verwenden dann halt eben mit "na".

   30:24
Mhm.
Den Cert Manager oder.

  30:27
Genau mit Let's Encrypt, dass du da den Set Manager mit denen einfach immer rollierst.

   30:28
Yeah.

  30:33
Oh, was.

   30:33
Da hast halt das Problem, das ist halt wieder irgendeine C L I, die du auf irgendeinem Host antriggern musst. So könntest du sozusagen deine deine Dings im Secret Store, wo deine Secrets hinkern, abbilden und sie dann einfach über den Agent zur Verfügung stellen.

  30:37
Gett.
Genau, genau und und.

   30:46
Das musst du halt abstimmen mit mit dem Rest von deiner von von deiner Infrastruktur beziehungsweise wie wird das in deiner Organisation hand gehabt, wenn wenn wenn es ein Team gibt, das das jetzt für die Macht, dann ist es natürlich Overhead. Andererseits, wenn du es nicht hast, ist das super praktisch.

   30:48
Ja.

  31:01
Ja, also das sehe ich schon, würde ich schon bei unserem Benefit zeigen, überhaupt wenn man dann schafft, interne PKIs anzubinden.

   31:02
Wunderbar, wunderbar.
Die Sache ist, sobald diese interne PKI es schafft, in den definierten Vault reinzuschreiben, hättest du das. Das heißt, dass sozusagen nicht wirklich ein Agent-Problem, sondern du musst irgendwie nur den Wert in den Vault reinkriegen und dann könnte der Agent schon seine Dinge betreiben, genau.

  31:16
Mhm.
Yeah.

   31:28
genau. Das also, ich glaube, wir sind uns alle einig, dass das konzeptionelle Modell mental gleich zu halten auf beiden Systemen dazu führt, führt, dass man schneller ist und halt effizienter arbeiten kann.

   31:43
Ah.

   31:43
Genau, seht ihr sozusagen diesen Git-Braced Approach, dass man den von meinem Agent gut in CICD Pipelines integrieren kann?

  31:54
Definitiv ja.

   31:57
Wunderbar, genau dann.
Gibt es irgendwas, was ihr sozusagen jetzt als Roadblock sehen würdet, um das nicht pilotieren zu können in der richtigen Entwicklungsumgebung?

  32:14
Würde mir jetzt nichts einfallen. Also, ich glaube, du bist eh schon ziemlich weit, deswegen ja.

   32:15
I guess.

   32:17
Mhm.

   32:18
Wollt Grad sagen, das fragst du jetzt, nachdem wir es auf Entwicklung rausgespielt haben, wie lustiger.

   32:22
Ja, das ist ja das aus Entwicklung ist ja nur ein Tom. Also glaube ich wisst ihr, was ich meine genau generell zu dem CICD mäßig. Das Schöne ist, es ist nur ein Curl Befehl, wo du ihm halt einfach ein Jammel mitgibst und entweder hast halt dann noch das Client CA.
Also, entweder musst du einen Request halt verschlüsseln und signieren mit M.T.L.S. mäßig und dem J.V.T. oder halt nicht. Genau, ja, dann sag ich einmal danke.
Stopp die Transkription.

 Transkription beendet
