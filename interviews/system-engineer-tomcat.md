Transkript
3. April 2026, 12:02PM

Transkription gestartet

  0:03
habe. Wunderbar, gut. Dann gehen wir mal rein. Du kennst dich ja eh schon ein bisschen aus, aber sozusagen gehen wir mal von Anfang an rein. Also, Secret Management in hybriden Systemen ist noch nicht ein ganz einfach triviales gelöstes Problem. Darum habe ich mich sozusagen mit meiner Masterarbeit damit auseinandergesetzt.
Dafür ist einmal der erste wichtige Stein zu wissen: Was ist ein Secret Store? Hast du diese Information oder sozusagen ein Centralized Secret Store oder soll ich das noch mal kurz erwähnen?

  0:30
Ja, passt. Ist OK so.

  0:32
Passt. Du weißt, was involved ist. Du weißt, warum es ihn gibt. Cool. Gut, der Extile Secret Operator ist dir ein Begriff und die Funktionalität. Sonst erkläre ich den noch einmal, weil der ist wichtig fürs ganze Konzept.

  0:45
Gehen wir es noch mal schnell durch, ne?

  0:47
Also in Kubanetis gibt es ja diese Secrets und die appliest du händisch oder du hast halt irgendetwas im Pot, der sich darum kümmert, dass diese Datei dann at runtime dem Container zur Verfügung gestellt wird. Der Exile Secret Operator gibt den Approach, dass er sozusagen ein Proxy ist, den man über

  0:50
Yep.

  1:03
Git passiert, dem man halt deklarativ sagen kann, was er zu tun hat und der stellt für einen halt nativ dann die Secrets in der OpenShift. Das heißt, dem sagst du: 'Red bitte mit dem Vault und gib bitte generiere mir dieses Secret aus den Daten aus dem Vault und der macht das für dich.'
Und dann kann seine Applikation ganz unabhängig von dem Ding mounten. Das heißt, wenn der mal ausfällt, kann die Applikation trotzdem hochfahren, insofern halt die Secret Values noch funktionieren. Und insofern du halt keine Ahnung, nicht alle Datenbanken mit 5 Minuten Credentials hast, dann ist es ja vielleicht okay, wenn der ausfällt, dass die Applikation trotzdem startet mit dem Passwort, das vielleicht
Die noch 10 Stunden gültig ist, soweit so klar.

  1:42
So.

  1:43
Genau. Jetzt stehen wir halt vor dem kleinen Problem, dass es sozusagen so einen einfachen Git-basierten, also Git-Ops-fähigen Operator nicht gibt für die Linux-Environments. Da gibt es ein paar Tools, ein paar Open Source, ein paar gebunden an Hersteller. Der, der am ehesten das könnte, was man brauchen würde, wäre der Hashikob Vault Agent.
Der unterstützt aber nur Hashikob Vaults, was sozusagen ja dann wieder gegen das System geht vom vom Excel Secret Operator, wo ich halt eine Vielzahl an Vaults und Auswahl habe als Backend. Genau, und demnach, nachdem es da nichts wirklich gibt, nach nach Literatur und Produktrecherche, habe ich mich dazu entschieden, das Ding selber zu bauen.
Und dazu habe ich mal konzeptionell den Workflow designed. Also, es gibt irgendwo ein Git Repository, da kann ich für meine Kubernetes Workloads sowie für meine Workloads, die auf Unmanaged Linux Hosts laufen, meine Konfigurationen einchecken. Und eine CICD Pipeline kann die dann applieren auf die diversesten Umgebungen.
Jetzt ist natürlich die Frage, was ist ein unmanaged Linux Host? Es gibt ja Linux Hosts, die laufen in einer Cloud, in einer Frame, die haben wir eben auf Education API, die Sie verwenden können.
Das haben dieses Glück, haben ja nicht alle VMs und sobald du sie halt on Prime hast oder Bare Metal oder irgendwo, hast du halt diesen Vorteil nimmer, wenn du nicht auf einem Hyperscaler bist. Das heißt, die Arbeit beschäftigt sich sozusagen in dem hybriden Szenario genau mit solchen Linux Hosts. Genau, und das Konzept ist, es gibt halt irgendwo eine Source of Truth.
Einen Sie eine CICD Pipeline und oder GitOps Operator, je nach dem, wie man es halt für sich selber konfigurieren möchte, und der spielt halt dann die Konfigurationen für den Secret Store und die Secret Configuration auf die einzelnen Systeme rüber. Wenn es irgendwas Unklarheiten gibt, dann funktionieren wir einfach dazwischen, würde ich sagen, oder?

  3:29
Du passt.

  3:31
Wunderbar, gut. Kommen wir zu meinem Solution Approach, den ich dir eh schon mal vorgestellt habe. Ich habe sozusagen jetzt im Zuge dieser Masterarbeit aufgrund eines Thread Models einen Agent gebaut, der sozusagen secure von Design sein sollte und sozusagen diese Aufgaben übernehmen.
Sollte die wir grad gesehen haben, was hier fehlt, und zwar ein System, was diese Config empfangen kann, die Dateien schreibt und die Dateien halt irgendwohin ablegt, dort wo sie erreichbar sind.
Genau, und das wird dann so ausschauen, dass du einen Secret Agent auf jeden Host, den du bespielen möchtest, drauf hast. Dem schickst du eine Konfiguration zum Store hin. Also, wie authentical er sich zum Store und wie und welche Secrets hättest du gerne erstellt und getemplated?
Genau, dann gibt es den Reconciler, der halt sozusagen einfach aufpasst, dass das Data da konfiguriert ist, auch wirklich am Host verfügbar ist und halt die Sequels wie Rolled. Dann haben wir den Atomic Filewriter, den ich dir eh schon mal erklärt habe. Da werden wir noch mal kurz ein Recap machen und es gibt halt den Metrics Endpoint mit Promiflash und Graphana, per se, weil das halt eine Open Source Applikation ist.
Man könnte deswegen wir das einsetzen, auch sicher mit *** und deren Go-Package.
monitoren, wenn man das möchte, wenn das so weit kommt. Genau. Also System ist, du hast den Pro Host, den du hast, einmal diesen Agent laufen und dem schmeißt du die Konfigurationen hin, wie du möchtest, was dort auf dem Host passiert. Und sozusagen fehlschlagende Rollierungen oder State-Sachen werden auf den Grafana Endpoint gepackt.
Published also ein Promifer is Endpoint gepublished und du könntest mit dem Stack dir sozusagen Alerts bauen, dass du sofort mitbekommst, wenn ein File Reroll nicht funktioniert, weil du würdest dir nicht also Fail Silent ist halt in so einer Architektur finde ich nicht wirklich das richtige System.
Weil du ja sonst nie weißt, ob deine Secrets richtig funktionieren und mit dem, dass du das halt scrapst über Promipheus, hast du halt auch ein wunderschönes Uptime Monitoring, ohne sozusagen viel extra in der Applikation einbauen zu müssen. Genau, kommen wir jetzt hier, genau dieser.

  5:18
Mhm.

  5:33
Dieses System frisst sozusagen 2 großartige Config-Typen und zwar einen Secret Store. Der hat, wie schon erwähnt, welcher Typ an Secret Store bin ich. Wie komme ich dort hin? In den meisten Fällen wird das eine URL sein, ein Authentication Block, wo ich halt meinen Weg zum Authentication Memvault reinschreibe.
Und sobald ich den applye, wird halt geschaut, ob die Connection so funktioniert, weil es wäre blöd, wenn du eine Connection ablegst, die du gar nicht hinbekommst. Genau, und dann gibt es Ja, dann gibt es noch Managed Files. Managed Files sind sozusagen so wie die Secrets, die wie die externen Secrets, also eine Config-Datei, die du an meinen Agen übergeben kannst.

  5:58
Funktion.
So.

  6:09
Arrest und da steht halt das Template vom File drinnen, von welchem von welchem Secret Store sich das holen soll und mit welchem Faden er was wie von holen darf. Genau, und das ist sozusagen eine Templating Auflösung in Go und es hat sowie das Secret Opera, Excel Secret Operator, eine Rotation Policy, das da halt auch dieses Dynamic Secret Management
bilden kannst und was er im Vergleich zum X. L. S. Grut Operator hat. Er hat auch eine JSON Jaml N. von PEM Validierung, der halt einfach schaut, ob das das Richtige ist, also ob das ob die das File per se O. K. ist, weil wenn du auch eindeutig für eine Applikation ein falsches Jaml hinlegst, ein falsch Intentit ist, fährt die Applikation immer hoch und das darf ja auch nicht passieren.
gehen.

  6:49
Genau, ja.

  6:50
Genau, desto das geht natürlich auch mit XML und so weiter. Aber bei XML hast du das Problem, wenn du kein Schema hast, kannst du halt trotzdem Fehler reinmachen, die ohne Schema nicht catchable sind. Die File Permissions haben sich in der Zwischenzeit auf 0644 geändert, schon also es gibt jetzt eine Gruppe.
An Prozessen, die lesen darf, aber das ist jetzt nicht so wichtig. Das ist eine kleine Feinheit. Genau, kurz einmal noch mal die Features zusammengefasst. Es gibt halt eine CRAD API für Store und für Secret Management. Das heißt, ich bekomme auch die konfigurierten Stores und Secrets zurück/managed files.
Aber nur redected, das heißt, ich kriege nur sozusagen die Konfiguration, die auch im Git Repository ist, zurückgespielt und beim Secret Store wird der ganze Aufblock halt nicht übertragen. Das heißt, die Secrets per se verlassen niemals den Host. Wenn Sie einmal drauf sind.
Und die Secret Values in den Dateien verlassen halt auch nie den Host. Dann gibt es einen Auditlog, wo jede Aktion, jede Rollierung protokolliert wird. Dann die Time-to-Liv, den Metrics Endpoint und sozusagen, man muss explizit in der Config des Agents angeben, in welchen unter Directories der schreiben darf.
Man könnte natürlich sagen, man löst das über den. Man hat das sowieso über den Linux System User abgelöst. Das kann man dann auch so konfigurieren, dass es ihm relativ egal ist, indem man halt Root gibt zum Schreiben. Aber wenn du zum Beispiel jetzt ein User hast, der ein bisschen mehr darf, oder du zum Beispiel willst, dass der Agent nur in Conf Directors reinschreiben darf.
Kannst du das halt so damit abbilden.
Genau, kommen wir jetzt noch zu ein paar Architekturentscheidungen, die sozusagen relevant sind. Ich weiß nicht, bei uns, aber per se sollte der Agent die Möglichkeit haben, auch Internet Facing Deploy zu sein. Demnach habe ich mich für einen Layout of Findication Approach entschieden, dass es sozusagen einen M.T.L.S. Check gibt.
und einen JVT Check. Das heißt, nur wenn es von einem validierten Client kommt und der JVT Pass, darf man weiter. Man kann natürlich auch konfigurieren, nur 1 davon zu verwenden. Würde ich halt abraten, wenn man ins Internet hängt, weil du sozusagen bei JVT CR die Replay.
Attacken hast, solange sie gültig sind und generell eigentlich niemand. Also generell finde ich neue Systeme ohne MTLS, wenn sie selber zu Server sind, nicht mehr konfiguriert werden sollten.
Genau, dann der Atomic Fileswap, wo er halt checkt, ob ich im Pfad drinnen bin, dass wir keine Practor versal vulnerabilities haben, wo er halt, wenn es eine Dateiformat ist, was erkennt und unterstützt, probiert er das zu marsheln. Und wenn er das marsheln kann, dann geht er OK. Wenn nicht, sagt er nein, da ist schon was.
Grob falsch, so kann das nicht sein. Dann kreieren wir halt den Temp File, dieses Temp File Permission. Also setzen wir die Permissions richtig. Dann schreiben wir das zu Disk. Dann syncen wir das und dann schließen wir dieses Temp File. Und erst wenn das alles funktioniert hat und wir wissen, dass sozusagen dieses File Host existiert.
Dann machen wir uns so eine Tomic Swap, der uns dann halt sozusagen die Gewissheit gibt, dass wenn wir eine Datei wirklich swappen, dass sozusagen die ganze Datei existiert hat und eine valide Datei ist und wir nicht so halbe Dateien schreiben, weil vielleicht der Storage und Server voll ist oder andere Use Cases Schrägstrich.
Einschränkungen über Container oder Sinems.
Genau, installieren kann man den Agent über ein Ensible Playbook. Das macht zur Zeit einen System User für den Agent, eine Secret Gruppe, wo sozusagen dann alle User rein können, sollten die sozusagen Secrets lesen dürfen. Darum auch 0644 Permissions.
Dann haben wir halt die Pfade, die halt ein Linux, die ihr halt auf Linux braucht. Dann templaten wir uns die Config zusammen, machen uns einen System-D-Service und das Binary wird am Host heruntergeladen. Genau.
Konfigurierbar ist relativ einfach über ein Jaml-File. Du könntest es auch über Environment-Variablen machen, aber ich finde es über beides unterstützt er. Ich finde es zum Herzeigen so schöner. In Trusted Root schreibst du halt rein, wo er überall reinschreiben darf. Also zum Beispiel bei uns: ***.
K slash.comcat minus Server, also minus App 1 slash.com zum Beispiel und dann App 2 slash.com und so weiter. Den Port, auf den du ihn erreichbar haben willst, die Logging Sachen. Wo hättest du gern den Audit File hin, dann den Pfad für deine lokale SQLite Database. Dann hast du halt die Security Sachen mit TLS.
Kannst du abdrehen, würde ich nicht empfehlen. Und hier kannst du sozusagen das Client CA File hinterlegen, auf das er zugreifen soll. All diese Sachen basieren darauf, dass diese Files schon am Host sind. Die Files könnte man sich sozusagen mit dem Agent dann konfigurieren zum Überschreiben, aber es gibt per se ein Bootstrap-Risiko.
Das heißt, die Sachen müssen irgendwie in irgendeinen SSH-based Approach, sei es Ensible, sei es rüberkopieren, davor schon am Server landen. Genau, dann haben wir auf Indication Mode, Non MTLS, GVT und Managed Identity. Der ist nicht wirklich relevant, der ist nur ein bisschen zum Testen implementiert worden.
Das heißt, ich könnte auch nur MTLS als Authentication verwenden, wenn einem das reicht. Aber der secureste Weg wäre hier das Certificate File hinzuzufügen, weil dann macht er einen MTLS-Check. Und wenn ich sie auf JVT schalte, checkt er noch den JVT. Da kann ich sozusagen mit Punktwellen known mit so einem IDP wie Key Clock oder Azure O aufarbeiten.
Aber ich könnte auch sozusagen einfach nur ein Scherz verwenden, wenn das mein Weg ist, wie ich mit Jewities arbeiten möchte.
Genau jetzt kommen wir auch schon zu Configuration Examples. Hier sehen wir sozusagen, wie es beim External Sequid Operator ausschaut.
Der die sind den Weg gegangen, dass sozusagen sie so machen mit Provider: Doppelpunkt, dann den Typen und die Tenant ID, Volt URL und dann halt hier die auf Indication Sachen. In Kubernetes hast du halt den Vorteil, dass du halt diese Secret API hast und diese Secrets halt.
davor noch mal anlegen kannst und hier referenzieren kannst, nachdem du das auf Linux jetzt nicht so einen einfachen System haben kannst, ist für die Prototyp Implementierung, für die Masterarbeit sozusagen, der ganze Authentication Block hier. Das könnte man aber natürlich dann mit Managed Identities oder mit anderen Dateiformaten Schrägstrich Authentication Methoden zum Auslesen
dieser Werte verändern. Dann könnte man nämlich diese Stores auch einchecken ins Git. Genau, mein Age unterstützt so wie Kubernetes, Jason und Yaml. Das ist hier, sehen wir jetzt Jason, in dem Fall aber das Gleiche könnte ich im Yaml genauso abbilden. Genau, dann haben wir sozusagen jetzt noch ein Secret Example.
Wir sehen hier, so schaut's beim External Secret Operator aus. Ich muss angeben, wohin geht das. Ich muss ihm einen Namen geben. Ich muss ihm sagen, dass es ein Template in dem Fall ist. Dann muss ich dem Pfeil nochmal einen Namen geben, weil du sozusagen in dem Kuba Neti Secret mehrere Dateien drinnen haben kannst.
Und dann kann ich hier halt mein ganz normales Kontext schreiben und dann habe ich hier ein Mapping von dem auf dem. Das heißt, er weiß, das hier ist lokal und remote brauche ich diesen Namen und das löst du dir dann halt auf in deiner Time-to-Live.
Genau, das Gleiche habe ich so circa ähnlich und ein bisschen versimplifiziert für den Linux Server, weil wir brauchen ja nur den Pfeilpfad und den Namen. Dann natürlich auch eine Referenz zu von wo wollen wir die Daten bekommen, unsere Rerolltime, Secret Report und was ich so sehr
Sehr balgy finde bei dem ist, dass ich so viel extra Text habe, um das Gleiche auszudrücken wie hier. Das heißt, bei mir ist das Mapping nur von hier auf 'Secrets' der Name, von hier auf 'Secrets' der Name.
Genau, das ist ja, das ist ja eine Liking-Frage und man könnte ja entweder einen Parser schreiben, der genau das auf das übersetzt, oder halt den Agent in Zukunft genau diesen Syntax auch unterstützen lassen. Ich hab mich mal für die Prototyp-Implementierung auf was Einfaches.
Zurückfallen lassen. Genau jetzt ist die Frage, willst du sozusagen kurz ein kleines Secret konfigurieren oder nicht?
Wenn nicht, skippen wir das, wenn schon.

  14:27
Doch, freut mich nicht, ne.

  14:29
Passt, dann skippen wir das. Das ist nur sozusagen für meine Kollegen, weil die dann damit arbeiten sollten und auch mit dem Excel Secret Operator, dass sie so ein bisschen Hands and Experience bekommen. Genau, kommen wir auch schon zum Form. Das ist in ein in Likert Scale aufgebaut.

  14:33
No.

  14:44
Bautes.
Aufgebaute Fragen.
Das heißt, die einfach kurz beantworten für mich. Wenn du Fragen hast zu den Fragen, gerne fragen. Darum ist es auch ein Survey und danach. Also dadurch ist auch ein Interview, ein Semi Structured.
Und danach haben wir noch, hätte ich noch ein paar Open Open Questions, wo ich dir ein paar stellen werde, wo mich deine, wo mich deine Expertenmeinung interessiert. Und dann habe ich noch ein paar Fragen zu dir als Person, also nicht Person, sondern deine Rolle, deine Erfahrung U. S. W.

  15:12
OK.

  15:18
Und das halt.
Dann in der Masterpad verwenden zu können, genau.

  15:23
Ach, also das sollte ich jetzt gleich ausführen oder?

  15:25
Ja, ja, ja, wäre am besten.

  15:26
Passt so, okay.
Du warst.

  18:47
Wunderbar, vielen Dank. Du bist, äh, es ist, soll ich sagen, was ich relativ interessant finde. Die Leute aus meinem Team haben mir bei den Fragen immer Rückfragen gestellt und irgendwie alle nicht aus meinem Team beantworten sie ohne mich was zu fragen. Also.

  18:58
He.

  19:02
ist nur eine interessante Beobachtung an der Seite. Genau, und wir wollten zu den Open Questions gehen. Per se läuft der Agent jetzt schon auf einen unserer Server und hat seit heute den ersten TomKit, den er bespielt. Er läuft bei uns gerade unter dem www.user halt.

  19:02
OK.

  19:20
Und ja, und per se legt er halt also die Datei, die du dann rüber kopierst, die sie ein. Also, was ich hier verstanden habe, das *** macht ja das, das nimmt die Dateien im Com-Folder und gibt sie irgendwie weiter dann an den echten Tom Kid, oder?

  19:21
Das passt.

  19:37
Der in dem Container läuft.

  19:40
Ist der den Kampf folder haben wir deswegen gemacht, damit uns die also dort haben wir die die Fuster Konfiguration und kopieren du das noch in den Suchordner Catalina Lockerlust.

  19:51
Genau, weil genau dort konfigurieren wir jetzt halt einfach die Datei rein und verwenden das zur Zeit. Das funktioniert. Es passiert jetzt noch nicht Security, Best Practice und so weiter, aber der Tomcat startet damit. Und wenn der und wenn der Rerollt halt ist und du den Tomcat dann startet, nachdem die Datei gererollt wurde.

  19:59
So.
Jo.

  20:09
Schaut wieder die neue Datei dar, gell? Genau, das heißt, wir sammeln einmal damit Erfahrung. Der Urtschild der Gemeinde möchte dann auch nochmal mit dir darüber reden, über das Ganze, was du davon hältst, wie du das Ganze machen möchtest in deinem Verfahren.

  20:17
Mhm.
Hook.

  20:22
Und ja, aber da gibt's dann hoffentlich in den nächsten Wochen ein paar Tage.

  20:25
Also, den den Dongkard Neustart hat es dann eher dabei, oder wenn du einfach die.

  20:29
Nein, den TomKat Neustart müsste man, muss man sozusagen noch händisch machen. Der ist nicht dabei, aus dem aus dem einfachen Grund jetzt einmal fürs Testen und Punkt und Punkt 2. Dieses TomKat neustarten kannst du nicht wirklich, wenn das File rolliert wurde. Weil wenn du das File, keine Ahnung, auf 40 Minuten einstellst, du kannst dir so etwa so große Applikationen

  20:35
OK.

  20:49
Nicht alle 40 Minuten neu starten, also wie man das sozusagen macht.

  20:52
Ne, das das brauchst eh nicht. Ja, einmal im Buckel ich die.

  20:54
Yeah.
Ja genau, und wie man das sozusagen genau macht, kann man ja dann noch ausdiskutieren, ob man einfach ein Signal rüberschickt von hier und so weiter. Wird sagen, per se für meine Masterarbeit ist es irrelevant, weil ich argumentier einfach, man kann einen Filewatcher auf dieses File haben, wenn.

  20:57
I'm our for you.
Yep.

  21:11
Einem interessiert.

  21:13
Mhm.

  21:14
Weil sozusagen, sonst mìsstet halt wieder dem Agen die Möglichkeit geben, dass er irgendwas mit der Applikation zu tun hat und das will ich ja nicht. Ich möchte, dass dieses Ding diese Dateien ablegt und danach nix mehr mit der Applikation zu tun hat.

  21:27
Na ja, gut, dann einen Ruck kommst vielleicht noch machen. Hast gesagt, man kann sich dort anmelden und dann ruft halt irgendwo auf, ja.

  21:29
Yeah.
Ja, oder man schaut halt einfach generell auf die Datei, die man braucht. Aber ja, das ist das funktioniert ja auch über Linux. Genau, gut, dann so.

  21:37
Ja.
Oder du startest einfach jeden Tag einmal neu, dann passt das auch.

  21:43
Ja, genau. Ja, wunderbar. Gut, dann kommen wir noch zu.

  21:47
Ja.
Das machen wir ja mit vielen Dank. Die starten einfach jeden Tag rein, weil es einfach.

  21:51
OK.

  21:56
Weil es einfach gescheiter ist, ja, weil irgendwer bastelt irgendwas zusammen, da bleibt dann ein Speicher irgendwo liegen, der nicht freigegeben wird, weil die Objekte nicht freigegeben werden und solche Sachen. Und wir haben die die Altlasten, die wachsen einfach.

  22:06
Verstehe gut.
Yeah, yeah.

  22:11
Und da mein Dock nicht start, du hast das Problem dann wieder.

  22:14
Genau da, genau deswegen hab ich mich halt auch für diesen Filebased Approach entschieden, anstatt sozusagen einen Agent, der probiert selber Sachen zu starten, weil du kannst mit rein deutlich jede Brownfield-Applikation anbinden, solange die ein File hat. Das heißt, wenn du irgendeine Software dazu kaufst in dein System, die sozusagen über Files.

  22:15
You.
Yeah.
Mhm.

  22:31
Die Secrets hat keine Ahnung, sagen wir jetzt, sagen wir, du kaufst dir Giro und betreibst es selber oder so oder diese Atlasian Suite, könntest es halt sozusagen so. Könnte man sozusagen auch mit so einem Proach anbinden. Genau, siehst du irgendwelche Vorteile in diesem Linux Secret Agent?

  22:39
Ja, das machen wir.
Yeah.

  22:49
Den ich dir vorgestellt habe / siehst du irgendwas, was dich stören würde an dem Konzept des Agents?

  22:57
Na ja, wann ich also ich bin davon ausgegangen, dass das anders läuft, nur ich hab mich damals wenig beschäftigt gehabt damit und ich tat es jetzt wahrscheinlich in Java machen.

  23:02
Yeah.
Wirklich.

  23:10
statt in Go. Ja, weil da haben wir die ganze Infrastruktur dazu und wenn du jetzt sagst, jetzt läuft es jetzt eh unter W. W. W. User, da hätten wir sogar schon 'n Dank gehabt, wo man die tollen Kunden, also das das war aber von dem her egal. Go ist auch gute Sprache, ja.

  23:21
Ja, ja.
Die Sache ist, ich habe das ja per se jetzt nicht geschrieben für unsere Firma, sondern generell einfach als Pardon zu dem External Secret Operator. Das heißt, ja.

  23:33
Das hast du gesagt, ja, was distanzöses, ja, ja, ist schon klar.

  23:35
Das war natürlich ja und das Schöne ist halt, du hast halt ein Safe Contain Binary. Das hast natürlich mit der Jar auch, aber du brauchst halt die JVM.

  23:45
Richtig, ja, ja, ja.

  23:47
Demnach war das.

  23:47
Bestimmt ja, aber aber ist wurscht. Ja, go, go, passt schon, ja.

  23:49
Ja.
Wunderbar, gut dann.

  23:53
Wir haben nur wenig Go-Light. Drum, das ist eigentlich die, ich glaub, ich bin noch eher da, ja.

  23:56
Yeah.
Ja, ich verstehe Sie. Ich weiß auch nicht, ob ihr ihn wirklich bei euch adaptieren könnt, aber vielleicht reicht meine Idee und meine Konzepte dafür, um das für euch richtig umzusetzen.

  24:03
jetzt
Ja, jetzt soll der der *** soll das jetzt einmal benutzen, weil der brauchts ja als Erstes.

  24:11
Yeah.

  24:13
Und da werden wir Erfahrung sammeln, sag ich jetzt einmal. Dann machen wir es noch was, das ist von dem her sowieso klar.

  24:16
Genau.
Yeah.
Würdest du sagen, der Security Tradeoff, dass du sozusagen Files schreibt, also dass der Agent Files schreibt, ist sozusagen okay im Vergleich dazu zu halt anderen Methoden oder würdest du meinen, das ist ein Dealbreaker für die meisten Systeme?
Well.

  24:37
Schreiben ist OK, weil das das kommt, das ist ja auch so eine Geschichte, dass er jetzt die Files dort hinschreibt, wo er sitzt, hinschreibt, das ist ja nur eine Lösung, ja. Man kann es ja anders auch machen, dass er es woanders hinschreibt und von dort dann geholt wird, ja.

  24:48
Ja.
Ja, ja, das geht natürlich auch, ja.

  24:54
Oder.
Das ist also keine Kalkibrika. Na, das passt schon.

  24:58
Verstehe ja, weil die Sache ist, was, was ich halt auch im Thread Model Acknowledged habe, wenn jemand halt die Disc der des Servers bekommt. Über was auch immer, es ist halt alles exposed. Das sind die Anmeldedaten.
Beim Beim Secret Stex Post für den Host und halt die Dateien, die auf dem Host liegen. Und sobald du root bist, hast du halt auch beides. Das sind halt 2 Security Concerns, die man erwähnen muss, genau siehst.

  25:20
Do.
Das stimmt ja, das in in Wahrheit haben wir das Problem ja sowieso.

  25:31
Yeah.

  25:31
Weil du musst irgendwo das Secret hinfahren, damit das der Agent auf Wolken treffen kann. Das muss er wissen.

  25:37
Genau, in Zukunft. in Zukunft gibt es ja dann Managed Identities und dafür gibt es ein Spy von Spyre, um das halt auch auf Unmanaged Linux Hosts zu bekommen. Du hast ja das Schöne in der Cloud, wenn du dort eine VM nimmst, hast du ja eine Off API, die du ansprechen kannst, die sozusagen deine Autorität sicherstellt.
Nur die Arbeit beschäftigt sich halt mit Linux-Systemen, die das nicht hat und für die ist das, finde ich, ein Tradeoff, den man gehen kann. Genau. Welches Szenario fällt dir ein Szenario?

  25:58
Ja.
Ja.
Und es ist ja, es ist ja noch was, die die selber laufen ja alle in den Container.

  26:10
No.

  26:10
Weiß nicht, ob er das der Ult schon mal erklärt hat. Das macht das ***.
Das macht der *** sprich, der macht ein eigenes Rootfeilsystem, wo er nachher nur die Sachen einhängt, die der Tank nicht wirklich braucht. Ja, das aus der Sicht.

  26:21
Ja.
Ja, ja, aber trotzdem eine W. W.

  26:27
Der Sicht die anderen Bank jetzt nicht, ja.

  26:30
Yeah.

  26:31
Obwohl die eben im selben Feld System liegen, aber der Sicht es nicht, weil es eben über spezielle Monts extra für *** zusammengestoppt wird, was er braucht.

  26:41
Versteh ja, da können.

  26:41
Und da sich da also auch, wenn du jetzt in ein Dunkelt einbrichst, du kommst dann nicht aus aus aus dem also die die Agency Clicks wirst du dann noch nicht finden, ja solche Loch.

  26:48
Ja.
Ja, ja, wie gesagt, das ist halt, wenn du wirklich einen Root-Compromise am Host bekommst, ist halt alles weg, aber.

  26:52
OK.
Ja, drum machen wir es ja mit Username Spaces.

  27:00
Yeah.

  27:01
die Container nicht mit, so wie der Boker mit der, das braucht man nicht.

  27:04
Ja, ja, ja, gut.
Siehst du irgendwelche Szenarien außer Application Secret Management, wo man diesen Approach von so einem Agent der Gitbase Konfiguration für dich in Reconsult einsetzen könnte, oder meinst du, Secret Management ist alles, was man damit abbilden kann?

  27:31
Ist.
Ob ich mir eigentlich keine Gedanken mache, aber momentan sehe ich da nichts.

  27:38
Verstehe, ein paar Kollegen haben gemeint, man könnte generell ganzes nicht Secret, sondern ganzes Configuration Management probieren, damit abzubilden, dass du sozusagen nicht mehr auf auf den Host passt, dass du halt nicht mehr auf den Host nur, also das sozusagen in Git deine

  27:38
Good much, Victoria.
Ja, ich seh Sie gleich. Ja, ja.

  27:57
Situation getemplated irgendwo einchecken kannst und das dann per Host aufgelöst werden könnte in.

  28:02
Ja, jetzt checkt Thomas jetzt in so eine Version.

  28:05
Ja, ja, aber mit den mit.

  28:05
Und vom von dem.
OK, also da sind nur die Templates eingeschickt. Bei dir jetzt, ne? Wir haben bei uns ja die ganzen Sequets auch mit drin, ne?

  28:14
Yeah, yeah, genau.
Ja genau, das das wird damit halt wegfallen. Du hättest dann halt einfach in dem Fall, also wie wir das bei uns umsetzen, rein theoretisch könntest du ja in in diesem Deploy Zyklus halt einfach den Agent dann aufrufen, lokal mit einem kleinen Skript.

  28:20
No.
Hm.

  28:33
Was dir das sozusagen dann übernimmt, dann hast du auch keine Internet ex Facing Dings, aber das ist halt bei uns ein spezieller Use Case und jetzt nicht so das, mit dem sich die Arbeit beschäftigt, genau.

  28:43
No.

  28:45
So dann.

  28:46
Das war ja eigentlich die erste Idee, die ich damals gehabt habe, wie das umgesetzt wird, dass man eben beim Deployment das zusammensucht und nachher.

  28:56
Yeah.

  28:56
Die Templates auflöst und und so das den Dumpgate konfiguriert, bevor gestartet, aber passt.

  29:00
Ja, ja, das so geht das wahrscheinlich in unserer Struktur wahrscheinlich auch gut und besser, sozusagen. Das dynamische Secret Management geht ja davon aus, dass du alle 30 Minuten deine Secrets rollen willst und halt Applikationen hast, die damit umgehen können.

  29:16
Ja, das haben wir nicht, ja solche Applikationen haben wir nicht, weil du müsstest dann den den Datenbanktreiber müsstest dann noch ändern zum Beispiel, dass du nicht neu starten musst solche Sachen ja.

  29:17
Ja, ja.
Ja, ja, du müsstest. Ja, du müsstest.
Ja, ja, genau. Aber darum, darum geht dieser Agent ja nicht über unsere Firma, sondern.

  29:27
Ja.
Lo

  29:33
Wie war das genau? Genau, das war's. Gut, das da hast du keine wirkliche Meinung. Das brauch ich nicht fragen. Siehst du irgendwelche fehlenden Features oder Schwächen im Design des Agents oder irgendwas mitbekommen, was dir auf der Zunge liegt?

  29:52
Aber ich nicht, aber das sicher wahrscheinlich dann erst, wenn ich mit damit gespielt habe, ja also.

  29:53
Steh.
Ja, ja.

  29:59
Aus dem Mim, ein bisschen kommt der Appetit, sagt Marlon, dann fällt mir da was ein.

  30:00
Ja.
Yeah.
Genau, genau, genau, genau. Was gibt es irgendwelche großartigen Sachen, die die aktuelle Organisation, in der du arbeitest, davon abhalten würden, auf so einen Workflow zu wechseln?

  30:18
Was? Keine Ahnung.

  30:21
Passt ja, wunderbar.

  30:23
Aber das, das glaub ich. Na, da, da gibt's kein Problem, ja.

  30:27
Verstehe.
Siehst du sozusagen den Overhead des Aufsetzens von dem ganzen Systems? Weil du musst ja die Agents rüberspielen, aufsetzen, konfigurieren sozusagen. EM als als Werfe dann sinnvoll im oder nicht sozusagen. Also, siehst du den Overhead, den du brauchst, um sozusagen damit

  30:46
Geht davor aus. Ich gehe davor aus, dass wir es automatisch sagen.

  30:47
Okay, ja,
Wunderbar. Ja, gut, dann sind

  30:52
Weil für das haben wir ein Siebel. Wir haben jetzt insgesamt, glaube ich, *** *** wo wir das potenziell einsetzen können.
Also, da müssen wir automatisieren, sonst geht es eh nicht.

  31:07
Verstehe, ja, wie gesagt, es gibt ein Ensible, was für uns läuft. Das ist gerade ein abgespecktes. Ich habe auch ein Ensible für größere Sachen in meinem Repository. Das Ganze ist per se Open Source. So, ich kann jetzt einmal die Transkription beenden, weil das ist alles, was ich für die Masterarbeit brauche.

  31:18
Ja.

  31:24
So.
Muss ich den Button noch treffen? Wirklich Teams.

Transkription beendet
