
Genau, gut, wunderbar, gut, dann weg.

   0:06
Jetzt wird er so, so, so Übung im Blöd melden, weil die letzten 5 Minuten so blöd gemeldet hab. Jetzt muss das aus.

)   0:12
Verstehe, verstehe. Wunderbar, gut. Also, dann beginnen wir mal kurz mit. Was ist deine aktuelle Position?

   0:21
Principle.
Du hast wieder Wischte.
Principle und Defox für die *** Services.
Um.
kümmere mich um
Den Betrieb der Applikationen in der *** für die ***.
Und ***.

)   0:43
Wie lautet denn diese Rolle?

   0:51
Oh, 19 Jahre als Defop und für das Linux-Umfeld seit 14.

)   1:02
Wunderbar, gut. Ha, bist du jeden Tag mit Applikations-, Konfigurations und Secretmanagement umgeben? Also, machst du damit etwas oder nicht?

   1:12
Ja.

)   1:14
Wunderbar! Bist du zufrieden mit den aktuellen Methoden, die er einsetzt?

   1:21
Größtenteils.

  1:21
Nix Falsches sagen.

   1:23
Größten Größtenteils, wobei ich immer noch Luft nach oben sehe und wir weiterhin dran arbeiten, um das Gap nach oben kleiner werden zu lassen.

)   1:35
Genau, gut. Hast du schon einmal mit dem External Secret Operator irgendwas gemacht und hast deine Meinung zu diesem Tool?

   1:45
Also bezogen auf diesen extern secrets, extern secrets operate and open shift ja, da hab ich was gemacht, ich find das Konzept sehr gut.
Hab's allerdings noch nicht allzu oft verwendet, leider, weil wir im Secret Store hängen, welcher freigegeben wird, aber ich hoffe, das wird sich in den nächsten Wochen ändern.

)   2:07
Versteh gut, dann das Gleiche noch mal bei dir, ***. Du bist.

  2:14
Es ist super, dass der *** das schon zuerst gemacht hat, weil dann habe ich die Frage schon mal gehört. Ich bin technischer Architekt für das *** der ***.

)   2:16
Ja, ich weiß.
Verstehe, wie lange machst du das Ganze schon?

  2:34
In der Firma bin ich seit 2011 und technischer Architekt, so jetzt offiziell als Titel, glaube ich, seit 2016 oder so was.

)   2:43
Verstehe.
Bist du täglich mit Applikations-, Konfigurations und Secret-Management beschäftigt?

  2:53
Ähm, vielleicht nicht unbedingt täglich, aber sehr häufig, ja.

)   2:58
Steh, bist du mit dem aktuellen Approach zufrieden?

  3:03
Mhm. Es ist O. K., aber wie der *** auch schon gesagt hat, bisschen Luft nach oben und daran arbeiten wir ja auch aktiv.

)   3:15
Versteh.

  3:15
Du vor allem.

)   3:18
Nein, ich war ich. Ich hab nur geforscht. Wir kennen uns gar nicht. Vergiss das nicht.

  3:22
Ach so, Entschuldigung, Entschuldigung.

)   3:24
Genau, gut. Kommen wir jetzt kurz zu einem kleinen Overview, das ja.

   3:28
Wer sind Sie? Wir haben Sie noch nie gehört.

  3:30
Wer sind Sie genau? Warum? Warum? Moment einmal, Social Engineering, das dürfen wir ja gar nicht.

)   3:37
Genau, gut. Also, kurzer Recap, den ich jedem gebe. Ein Centralized Secrets Store ist euch beiden ein Konzept. Das ist ein Hard and System, in dem ich den den Life Cycle von Secrets managen kann und von wo Secrets retrieved werden können.
Soweit so klar, das Ding, was das macht und warum man das verwenden möchte.

  3:59
Mhm.

)   4:01
Wunderbar.

   4:01
Ja, aber ich kenne CSS und anderen Begriffe auch noch.

  4:04
Mhm.

)   4:04
Ja, genau. Gut, dann kommen wir auch schon mal zu meinem Proposed-Workflow für.

  4:10
Ja, wir müssen das ernsthafter machen. ***, Entschuldigung.

)   4:13
Nein, nein, du könntest das ruhig gern genauso weitermachen. Genau, also der mein Proposed Workflow ist, wenn man sich, nachdem ihr beide schon das Konzept vom External Secret Operator erkennt, das Konzept ist, dass es ein Git Repository gibt, was Secret Configurations für für Kubernetes.

   4:13
Tschuldigung.

)   4:29
Und für Linux Hosts hält, denn eine CICD Pipeline plat diese Changes dann, wenn sie im Main kommen. Und in der Componatis Welt macht es halt der Secret Operator für uns. Und in der Linux Welt macht es halt ein Secret Agent.
Den ich gebaut habe und das Wichtige ist bei beiden ist, dass die sozusagen Secret Store Agnostics sind. Das heißt, die unterstützen mehr als nur einen Anbieter im Secret Store und untersitzen auch pro Installation die Möglichkeit, mehrere Secret Stores zu verwenden, was nicht bei allen der Fall ist.
Sagen, dass die der also die Konzeption des hybriden Workflows habt ihr verstanden, würd ich mal sagen.

  5:09
Mhm.

)   5:10
Alles passiert per se auf in der communities Welt. Secrets werden generiert und in der Linux-Welt Files werden generiert.
Mit allen Pros und Cons, nur die du davon hast, wenn du eine SQL-Datei irgendwo herumliegen hast.

   5:20
Mhm.

)   5:25
Aber sozusagen System ist du hast.

  5:27
Da wird man nicht umhinkommen, befürchte ich.

   5:29
Genau, irgendwer muss das Sigrid lesen können.

)   5:32
Yeah.

   5:32
Und wenn du keine Secret-File herumliegen hast, dann hast du ein Secret herumliegen, das was dann von dem Server aus irgendwo auf den Secret Store hinkommt. Somit hast du das gleiche Problem wieder.

)   5:43
Es gibt auch andere Approaches per se mit Workload Identity und so weiter, aber für das, was ich mir anschaue, für das Managen dafür, ist das ein Security Tradeoff, den ich einfach hier von Anfang an acknowledge möchte. Genau, dann schauen wir uns kurz an, was macht Externe Secret Operator jetzt einmal genauer, um da ein bisschen mehr Verständnis reinzubekommen. Es gibt einen Volt.

  5:43
Mhm.
Mhm.

   5:55
Yep.

)   6:01
Von dort werden Secrets geholt und daraus werden Kubernetes native Secrets gemacht und die kann der Applikation mounten. Was wir hier gleich sehen, wir haben hier einen kleinen Proxy Adapter, was auch immer, der die Kommunikation übernimmt und den kann ich halt Git Ops basierend mit Kubernetes.
Resource Definition konfigurieren.
Und das Schöne ist, dadurch, dass sozusagen immer ein Secret da ist, kann eine Applikation auch starten, wenn der Secret Storm nicht erreichbar ist, insofern die Werte halt noch funktionieren zur Datenbank und so weiter. Aber per se bist du nicht runt and dependent auf diesen Agent.

  6:33
Mhm.

)   6:37
Und sozusagen, es gibt zwei große Konfigurationen, ja.

  6:38
Kannst kurz noch mal zurück.
In der Überschrift steht EOS statt ESU.

)   6:45
Dankeschön, werde ich dann gleich anpassen. Genau, also ist jetzt ein.

   6:49
End of support.

)   6:54
Die Sache ist, es gibt 2 Konzepte für das Mental Model, was wir brauchen. Und zwar, das ist die Secret Store Config. Das definiert den Zugang und die Offentication zu einem Secret Store und ist sozusagen dafür da, um halt Secret Configuration von den Offentication Details.

   6:54
Hmm.

)   7:10
Zu splitten, so könnte man das reinisch in 2 Repositories oder halt so aufteilen, dass Developer niemals die Authentication credentials sehen.
Genauso wie wir dann ein Managed File haben oder auch Kubernetes Secret oder External Secret genannt. Das ist ein Templating Approach, wo ich mit einer Referenz zu einem Secret und einem Secret Store und einer Time to Lift Policy dafür sorgen kann, dass diese Datei rerolled wird.
Genau, die beiden Sachen sind einmal klar und warum die gebrauchte Nexal Secret Operator umfällt.
passt. Damit propose ich sozusagen für diesen holistischen Workflow, den ich schon vorher erklärt habe, einen External Secret Operator für Linux Systeme, der sozusagen die gleichen Konzepte eines Store, einer Store Config und eines
Templates für Files hat und das halt für einen übernimmt, um halt sozusagen mental model mäßig genau gleich unterwegs zu sein.
Genau hier, noch einmal, noch einmal.
Hier noch einmal das Ganze, aber das haben wir ja alle, glaube ich, relativ gut verstanden und haben da schon Vorerfahrung. Das heißt, da brauche ich jetzt nicht noch einmal tief reingehen. Genau, wie wird das denn sozusagen ausschauen? Dafür ist ein schematisches Bild. Ist was wäre, wenn dieser Agent auf 3 Servern läuft? Dieser Agent hat eine lokale SQLite Database, in der er seine
ganzen Configs und auch die Connection Strings reinschreibt zu den Secret Stores. Der hat eine API, über die er aufrufbar ist und er reconsult das, das heißt, die konfigurierte Time-to-Leaf wird eingehalten und du hast einen Slash Metrics Endpoint, den du mit Promifoles Verfahren und Alert Manager so konfigurieren kannst,
Dass du alle Daten über Uptime hast und wenn ein Secret Mal nicht rerollen kann, dass du Alerts schreiben kannst, dass du sozusagen kein Fail Silent hast und auch sozusagen nicht nicht gezwungen bist, ständig zu schauen, ob alles funktioniert. Genau.
Das heißt, das ist ein Binary, was konfigurierte Pfade hat, wo es Files hinschreiben darf. Diese Files muss ich über eine API hinkonfigurieren und den Rest macht dann der Agent.
Solution Konzept verstanden, würde ich mal annehmen, nachdem ihr beide so still seid.
Ja.

  9:23
Schweigen heißt Zustimmung.

)   9:24
Genau, reden gehen wir kurz.

   9:25
Oder wir sind halt fertig mit Denken.

)   9:29
Gehen wir kurz mal auf die Features vom Agentan. Es gibt eine Crat API für Secret Management, auch einen Get Endpoint. Und da beim Get Endpoint kriege ich sozusagen meine konfigurierten Files und Distores, aber mit Redacted Information zurück. Das heißt, ich habe auch eine Möglichkeit.
Mir selber drift.
Drift zwischen meinem CI Repo und der echten Welt hier anzeigen zu lassen. Das Ding schreibt einen Audit Log, wo jede Re Roll, jede Änderung dokumentiert wird. Jedes Secret hat eine Time Drift und wird rerolled, den Metrics Endpoint und jeder Pfad. Darf ich denn das E vergessen? Ich will auch noch gleich ändern.
Und jeder Pfad, wo das Ding hinschreiben darf, muss explizit erlaubt werden.

  10:11
DTL Reroll, meinst du, nach dieser Zeit musste Secret spätestens neu aus dem Secret Store geladen werden?

)   10:11
Genau.
Nach genau, nach der Zeit wird es revolt, genau. Aber dazu wir kommen später dann noch beim Konfig Example, was ihr beide auch anlegen dürft und machen müsst. Genau nach Ja, nachdem sozusagen dieser Agent.

  10:20
Mhm.
Yeah.
Jetzt müssen wir tun, auch noch was.

   10:29
Ja, das das ist die schlechte Nachholzimmer.

  10:32
haha

)   10:32
Nach dem der Agent sozusagen dafür konfiguriert ist, ist also ausgelegt, ist, dass ich ihn auf jeglichen Server schmeißen kann oder auf jeden Linux Server. Der und ihn im Internet exposen kann, ist die Authentication Architecture, die man verwenden kann, ein MTLS Check und dann ein JVT Validator.
Was sozusagen das Ganze relativ sicher macht, weil zum Beispiel nur die CI dürfte mit dem Ding reden über halt den MTLS und du brauchst trotzdem noch den validen JVT, der eine kurze Expire hat und JVT, wie sie beide, kann ich mit einem Public Key ganz einfach validieren. Das heißt, per se muss ich da nicht unbedingt mehr Secret Infos dem Agent
geben, als er eh schon hat. Und ja, das Einzige, was du halt hast, ist ein Boost Rep Problem von den ganzen TLS Sachen. Aber das hast du leider bei allen TLS Systemen, dass du irgendwie die TLS Secrets dorthin bringen musst oder du generierst sie erst am Host.
Genau, aber sozusagen die Aufarchitecture, die möglich ist, ist klar, würde ich sagen. Du kannst MTLS konfigurieren und Javity.
Genau das Dings swappt die Files atomically und nachdem das sozusagen dann schaut halt, ob er das File platzieren darf und nur wenn es genug Platz hat, dann macht er ein Temp File, dann renamed er die, weil das die einzige Atomic Fileswap Operation auf Linux ist. Und er schaut auch, ob wenn sozusagen der Content XML oder Jason oder Jammerl ist, schaut ob es überhaupt ein valides.
Jammerlich, weil vielleicht wurde ja ein Secret falsch eingezeichnet, also falsch eingegeben oder falsch escaped. Dann würde uns das auffallen. Wir würden keine kaputte Datei irgendwo ablegen und wir würden dadurch sozusagen das als Fehler markieren und dann würde sozusagen im Alertmanager ein Alert aufpoppen, dass da was nicht funktioniert.
Und wir verifizieren sozusagen den Pfad, wo wir hingürfen. Man könnte natürlich sagen, das kann auch einfach der System User, aber das ist sozusagen ein extra Security Feature, um halt einfach hundertprozentig sicher zu sein, dass Paftraversals nicht funktionieren, auch wenn der User, der den Agent ausführt, vielleicht nicht perfekt konfiguriert wurde.
Genau, der Agent wird installiert mit einem Ensible Playbook. Das ist einfach ein Install.yaml. Das macht einen Systems User, das kreiert eine secret Gruppe und das kreiert die Pfade, die eine Linux Applikation braucht. Das Template dir die Konfiguration.
Create dir das System T Service und enables dir und es lädt das Binary von meinem GitHub runter.
Genau, die Konfiguration von dem Agent ist relativ einfach. Ich habe eine Agent Kategorie, da sage ich halt, wie es heißt. Dann habe ich hier meine Liste an Trusted Roots. Das heißt, da gebe ich halt ein, welche Pfade dürfen beschrieben werden. Dann halt Server Settings wieder Port, dann die Logging Settings, wohin das Audit File geschrieben werden soll.
Dann der Database Path, also wohin er das Database File schreiben soll, sowie die Security Section, wo ich TLS halt auf und abdrehen kann mit dem CERT und den Key File. Und wenn ich MTLS will, kann ich das Client CA File hinhängen. Für Authentication haben wir None MTLS und JVT, das heißt, je nachdem, was ich möchte.
Wenn ich JVT und MTLS möchte, setze ich das Client C. A. File. Dadurch ist es mal MTLS sicher, aber die User auf Indication wird dann über den JVT gelöst. Und hier kann ich das entscheiden, ob ich das über einen Identity Provider lösen möchte mit den URLs oder ob ich selber das Secret.
Das Schert Sigrid mitgeben möchte.
Gut, ich glaube, Konfiguration ist auch relativ straightforward.
Kommen wir jetzt auch schon zu Configuration Examples. Wir haben den Asia Store jetzt auf bei den Systemen. Das hier ist der XL Secret und das ist das, was ich gemacht habe. Er unterstützt Jamel und JSON sowie Kubernetes. In dem Fall ist das jetzt JSON, das heißt.
Hier, der Asia Keyword schaut bei ihm so aus. Das ist schaut bei external Secret Operator so aus, dass du halt hier eine Tenant ID hast. Die habe ich hier im Authentication Objekt abgebildet und die Client ID und das Client Secret. Das braucht ja keiner sehen und du brauchst deinen Namen.
Wie das Ding heißen soll, dass es referenziert werden kann und du gibst einen Typ an. Die der Excel SECL Operator ist so gelöst, dass du den Typen hier als Key angibst und halt die ganzen Sachen so konfigurierst. Aber ich glaube, es ist relativ klar, wie ich einen neuen Asia Key wollte.
in meinem System und im anderen System anlegen würde mit dem Screenshot.
Genau, kommen wir jetzt zu: Wie konfiguriere ich jetzt ein Secret? Wenn ich diesen Store hinzugefügt habe, dann sehen wir hier die Kubernetes Weise. Der hat halt ein paar mehr Features, demnach ein paar mehr Fields. Das heißt, aber im Großen und Ganzen ist es gleich. Ich habe hier ein Name zum Secret Store, wie erwähnt, so wie hier.
Ich habe eine Rerolltime oder hier Refresh Interval genannt.
Beiden Fall 25 Sekunden. Ich habe hier eine Datei.
Hiermit dem Content. Da kann ich sozusagen das dann schreiben und hier verwende ich das Templating. Der XL Secret Operator mappt das auf das, auf das, das heißt, das heißt hier in unserem Template DB User.
Und heißt im Secret Store DB Username.
Bei mir, ich hab das ein bisschen einfacher abgebildet, indem ich gesagt habe, dieses Mapping von hier, das muss auf das passen und das ist die Referenz im Secret Store.
Soweit so klar, wie ein Template funktioniert im Config Dings.

  15:49
Ja, ich habe da sowohl im einen als auch im anderen Beispiel das Problem: Was ist, wenn im Buzzword ein Anführungszeichen vorkommt?

)   15:58
Wenn den Part dann würdest sozusagen im Kubernetes Dings würde wahrscheinlich das Anführungszeichen einfach machen und die Datei wäre kaputt. Bei mir würdest du einen Fehler bekommen aufgrund aufgrund des Valid Contents, weil ich sozusagen, wenn es ein Chase, ein Yummle oder ein XML ist.

  16:07
Ah, OK.

)   16:13
Validiert er per se, ob das also er probiert es zu marscheln und wenn nur wenn er es marscheln kann, macht er weiter. Das ist ein Problem mit seinem Secret Operator, was er.

  16:13
Mhm.
Mhm, obsolidays.
Uh.

   16:20
Das ist so, das ist so alles, was man im Doncat quasi unter Anführungsstriche auch haben. Nur da haben wir es mit Velocity unter Anführungsstriche gelöst, weil wir da das Escape XML machen.

  16:30
Yeah.

   16:31
Und dann tu ich das mit der 1 versenden, äh.

)   16:35
Genau, aber auf dem Linux-Dings kann das nicht passieren, insofern du beim XML das Richtige hast. Nein, also nein.

   16:38
Yo.

  16:41
Dann kann es schon, kann es schon, doch du kannst im Passwort da Anführungszeichen haben und dann irgendwas A. ist gleich und dann wieder Anführungszeichen, dann geht es wieder weiter. Und das ist aber lides XML, aber es ist trotzdem falsch.

   16:42
Kantone.

)   16:44
Then well's party, yeah.
Ja, aber dann.
OK, ja, per se könnte man dann für den Use Case per se halt noch ein richtiges Schema für das XML auch anlegen.

   16:56
Mhm.

  17:01
Na, du müsstest, du müsstest ihm eigentlich sagen, ob und wie er es escapen soll, ob er nach XML-Schema escapen muss, ob er nach Jamel escape muss, ob er nach Jason escapen muss, aber.

)   17:09
Ja, ja.
Das Das macht er mit der Dateiendung per se nach was es aber aber ja.

  17:16
Ah, er tut es geben.

)   17:18
Ja, also ich bin mir nicht ganz sicher, ob wir daran vorbeireden, aber das ist sozusagen ein ein kleines Detail, was ja sozusagen User Error in beiden Systemen ist.

  17:26
Mhm, ja, also anders gesagt, man muss schauen, dass man nicht irgendwie Inject-Probleme bekommt.

   17:33
Ja, erstens dies und zweitens, geil, muss das wirklich mit X.M.L. escapen, das Passwort, weil du musst ja kaufmännisches und escapen, weil sonst das beim Hochfahren von der Applikation ja gar nicht kennen würde zum Teil.

  17:37
Mhm.

)   17:45
Aber das ist ja per se im Secret Store ein Problem. Aber ja, ist acknowledged, werde ich aufführen in den Limitations.

  17:45
Genau.

   17:51
Im Secret Store vielleicht nicht, weil im Secret Store war sie mit ungespeichert. Aber das das war jetzt auch ein Problem, was man beim wahrscheinlich der Secret Operator nicht können würde. Der hat nur den Vorteil, dass dort ein

  18:00
External secret.

)   18:01
text

   18:07
Ding hast.
Jamel Host und da vielleicht mit dem Escape mit so ein Problem hast oder Anführungsstriche, war da trotzdem ein Problem beziehungsweise bei Jamel ein Problem ist, ist das Raute. Was das nicht unter Anführungsstriche, also was das nicht unter Anführungsstriche seht so aus dem Rauteproblem, weil da noch ein Kommentar hinter bei hat.

)   18:20
Ja.

   18:25
B. Z. Wonst du das mit Anführungsstrichen, du siehst, du hast mit Anführungsstrichen kein Problem.

  18:28
Mhm.

)   18:28
Ja, also.

   18:29
Oder du hast immer nur den Vorteil, dass du bei der Jamel und beim Sigrid hergekostet und kostet das Base 64 Kodieren noch hast, da kann man Thema nicht.

  18:38
Also das so was glaub ich muss muss man sich dann, bevor man es produktiv macht, noch anschauen, aber konzeptmäßig verstanden und OK.

)   18:38
Ja, genau.
Yeah.

   18:45
Genau.
Yeah, yeah.

)   18:47
Wunderbar, gut, dann kriegt ihr auch schon 'ne Aufgabe. Ihr müsst jetzt ein Secret konfigurieren, ich schick euch jetzt dann gleich in den Chattermin 2 Rohlinge 2 Rohlinge an und was ich gern hätte, wär ein Jason Secret, was ein Objekt ist aus Secret und das Value.

  18:51
Hehe.

)   19:04
Soll im Secret Store auf Secret Dev hinzeigen.

  19:07
Ah, das ist ja schrecklich.

)   19:10
Mhm.
Und FilePath könnt ihr wollen. Wie könnt ihr wählen, wie ihr wollt.
Das ist einfach ein.

   19:16
Du hast wieder ESO geschrieben.

)   19:18
Ja, das ist ein ist leider ein Folgefehler bei mir, da hast du recht.

   19:19
Mir gut aufgefallen, ja.

  19:21
im

)   19:25
Genau hier habt ihr sozusagen den Rohling, den ihr gerade auch in der Präsentation gesehen habt, und daraus baut bitte.
Einmal dieses Secret Content mit den Anforderungen, die ich euch hier gebe. Das ist einfach nur dafür da, dass ihr das auch einmal gemacht habt und nicht nur gehört habt, sondern ein bisschen ein Gefühl habt, für wie es ist, hier ein Secret zu erstellen.

  19:45
kriegt man nachher auch Prüfung.

)   19:47
Nein, aber ihr habt danach ein Form auszufüllen und dieses Form.
Ist halt, argumentiere ich, ist wahrer und besser, wenn man schon einmal mit dem Ding etwas gemacht hat.

  20:00
Mhm.

   20:08
Bist du also optimistisch mit einer 5-Runden-Hinweisung oder mit einer kurzen Einweisung, dass man dann gleich damit arbeitet?

)   20:15
Mhm.

   20:16
Setzen Sie jetzt gewaltig unter Druck, ich sag das nicht.

)   20:19
Ich.
Ja, wenn ihr Hilfe braucht, sagt das einfach. Dann helfe ich euch. Es geht nicht, es ist keine Prüfung per se, es geht einfach nur.
Darum, dass ihr das einmal selber konfiguriert habt und dann kann ich euch zeigen, wie einfach das App bleibbar wäre am Server.

   20:38
Du bist jetzt wirklich auf Server, das was wir verwenden würde.

)   20:40
Nein, also ihr macht das, ich schau mir das an und sag mir, wie cool, ob ihr das richtig gemacht habt und dann zeig ich euch, nein, nicht am Server, also ihr macht es einfach, dann zeigt es mir, weil am Server applian bringt uns ja relativ wenig. Also

   20:45
Ja, am am Server war.
Wo, wo, wo, wo müsstest du das jetzt machen?

)   20:56
Einfach in eurem Texteditor.

  20:57
Im Notepad plus plus oder?

)   20:59
Ja, irgendwo einfach dieses Secret konfigurieren mit dem Content und den Einstellungen.

   21:04
Ja, was du uns eh schon geschickt hast, haben wir uns eh schon.

)   21:06
Ja, aber das ist nicht das Richtige. Das ist nur ein Example.
Das ist nur, dass ihr nicht alles von Grund ab schreiben müsst.

   21:17
So, du musst dafür erstmal für ein irgendein bestehendes Kontext XML das einfach mal machen.

)   21:22
Nein, für nein, nein, der Content soll nur das sein. Der Content soll das sein.

   21:26
A true.

)   21:28
Ultra einfach nur, dass man mal ein Gefühl dafür bekommt, was man zu tun hat.

  21:41
Also praktisch nur die Werte ausbessern oder umändern.

)   21:43
Genau, nur die, genau, nur die Werte semantisch ausbessern, dass das halt übereinander stimmt, dass ihr halt sagen könnt, ihr habt einmal den Secret mit dem gemacht.
A true.
Und am Server Applierbar ist das Ganze mit einem ganz einfachen Curl-Befehl. Also das, was man in das EI integrieren müsste, wäre ein ganz normaler Curl-Befehl, der mit dem Jammel.
Die Daten nimmt und sie halt einfach erstellt.

  22:15
Mhm.

)   22:16
Das ist alles.

  22:17
Kannst du mal zurìck bitte?

)   22:18
Ja, ja.

  22:23
OK.
So schätz ich mal in etwa, oder?

)   23:03
Genau Content Punkt my secret my secret secret DEF wunderbar sieht man.

  23:11
Streber ich was?

)   23:11
Du hast.
Du hast verstanden, ja.

   23:13
Ja, warte, das weißt du nur ein bisschen schneller.

)   23:16
Ja gut, du hast doch erst 10 Minuten später begonnen mit deiner Arbeit.

   23:17
Right, jetzt muss ich diese.
Ja, aber hier, wenn du mit den ***** musst, gell.
Gib ihn immer zum zum zum Reviewen hättest das ja gemeint.
Aber ich glaube, ich schau ziemlich gleich aus.
Okay, ich hab sie, ich hab sie reingerückt.

)   23:37
Ja, es ist wurscht.

   23:37
Hey, wieso hast du das auch mein Sigrid genannt? Das ist ja gemein.

  23:49
Wir arbeiten uns halt schon lange genug zusammen, ***.

   23:51
Hey, cover.

)   23:54
Wunderbar, hat bei euch bei Ja.

   23:55
2 Idioten, ein Gedanke, sagt meine Frau immer.

  23:57
Mhm.
Mhm.

)   24:00
Wunderbar, wunderbar. OK, Ursch, du hast es in Sekunden gelöst. Das Sigma hat den er unterstützt sozusagen eine Stunde. Das heißt, du musst es nicht in Sekunden lösen.

   24:11
Okay, das das war ich mir nicht sicher. Deswegen, ich geh auf Nummer sicher. Ich hab never touch a running system.

)   24:12
Ja, ja, aber genau das, ja.
Genau, hier nochmal das Konfig-Example und jetzt kommen wir auch schon dazu, dass ich euch den Link zum Form schicke. Die Fragen jetzt dann einfach durch beantworten. Das ist Likeard Scale Evaluierung, also ihr müsst halt immer auswählen, wie toll.
Wenn ihr eine Frage nicht ganz versteht oder Kontext zu einer Frage wollt, fragt ihr mich. Das ist der Grund, warum das ein Interview ist und nicht einfach nur eine Form.
Es ist nur ein cooler Weg für mich, ein bisschen quantifizierbarere Daten über das Ganze einzuholen. Danach haben wir dann noch ein paar offene Fragen, zu denen, die ihr nicht alle beantworten müsst. So, wenn ihr irgendwo einen coolen Inset habt, freue ich mich natürlich dann darüber.

  24:48
Mhm.

)   24:58
Genau.

   25:02
Möchtest du mir das gleich jetzt beantworten oder?

)   25:04
Ja, ja, wär am besten, wenn es gut.

  25:06
The purpose and functionality of the Linux Secret Agent are easy to understand für für mich oder für für ihren Jemen OK.

)   25:13
für dich, für dich, du antwortest aus deiner Person, wie du, wie du das verstehst.

  25:18
OK.
Mhm.

   26:02
Manchmal ist es ein bisschen schwierig, wenn man noch nicht damit gearbeitet hat. Nein, bin ich. Bist du schon bei der Letzten?

  26:03
Die letzte Frage ist schwieriger.
Ah, also die, die im im im ersten Dings die Jensible Based Installation so eine Approvia Deployment Method.
Weiß nicht, vielleicht wahrscheinlich haben wir Enzeplin in Benutzung.

)   26:24
Ja, es stimmt.

  26:24
Keine Ahnung.

   26:24
Ja.
Oder hätte ich gesagt, normal, weil das ist ja, das ist einfach so.
Die, die, die, die ist gemein für die, die Frage, weil N sie wissen so nicht gut bei der Verwendung, das ist ja.

)   26:39
Ich mag Gänsebär.

   26:42
I only two with bones is.

)   26:45
Doch, es ist bei euch jetzt auch nicht so schlimm. Ich habe mich da reingearbeitet. Das ist halt ein bisschen eine Umstellung, aber.
Du musst halt Jammeln mögen, so wie ich.

  27:00
Hehe.
separation of secret store connection and secrets to be retrieved, provide security improvements. I wäre ja fast geneigt, da Nein zu sagen, weil Mhm.

   27:40
Ja, da häng ich auch gut.

  27:54
Weil ich glaube, wenn man jetzt rein nach Sicherheit geht, wäre es wahrscheinlich sogar geschickter, wenn die Applikation selbst zum Secret Store hingeht. Weil dann kann er Secret Store auditieren, die Applikation hat er zu dem und dem Zeitpunkt das und das Secret abgeholt.

)   28:05
Na ja, aber es gibt auch.
Das könntest du rein theoretisch ja so auch abbilden.

   28:14
Mhm.

)   28:17
Indem du halt einen Secret Store pro Application hättest.

  28:21
Mhm.

)   28:21
Dann wär das sozusagen vom Agent auch so abbildbar, rein theoretisch von

  28:25
Ja, aber der Agent holt sie ja immer ab.

)   28:29
Der Agent holt mit der Secret Store Connection ab, die du ihm gibst. Das heißt, wenn du dir eine Secret Store Connection nur für die eine Applikation machst, dann könntest du es eindeutig so auch abbilden.

  28:37
Mhm.
Das ist schon, aber.
Done.
Wenn es jetzt rein um Sicherheit ginge.

)   28:51
Ja, aber die Sache ist, das Problem ist immer, wie kriegst du die Credentials zum Secret Store in die Application rein?

  28:58
Genau, wenn die Application die Sigrid selbst abholen würde.

)   29:02
Ja, aber dafür muss, dafür müsste sich die Application zum Secret Store authentifizieren, authenticated und wohin gibst du, wohin gibst du retical dentions?

  29:02
dann hättest du dies den
Genau, genau. Ja, OK, das Problem hast du immer noch. Ja, das stimmt, aber das.

)   29:13
Dafür brauchst du halt Managed Identities und Workload und Workload Federal.

  29:16
Genau, genau. Wenn du das hast, dann wäre es wahrscheinlich auch sicherer, wenn du es anders machst. Deswegen.

)   29:20
Genau, aber.
Aber.
Ja, du, es ist deine Meinung als Experte hier gefragt.

   29:35
Now this is the gleichen problem hades by external secret operator you are by OpenShift.

  29:39
Definitiv, ja, ja. Hast das gleiche Problem. Also, für den, wenn es nach Sicherheit geht, ist es das, wenn es nach Laufzeitstabilität geht, im Sinne von, ich bin nicht so laufzeitabhängig, dass der Secret Store da ist und.

   29:52
Mhm.

  29:53
ich tue mich leichter beim beim Passwörter ausrollen und ich muss die Applikation nicht angreifen. Also das hat halt andere Vorteile, aber aus secret, also aus aus Security-Sicht, glaube ich, wäre es wäre es sinnvoller, wenn es die Applikation abholt. Die Frage ist jetzt, wenn du da schreibst,

   30:03
Gibt es schon Rechte?

  30:12
Provide Security improvements im Vergleich zu was?

)   30:15
Nachdem sozusagen die Lage bei jedem eine andere ist, kann ich das nicht quantifizieren.

  30:21
Ja, das verstehe ich. Aber dann ist eben die Frage, wie ich die Frage beantworten soll. Deswegen nehme ich es neutral.

)   30:25
Ja, ja.
Verstehst ja, du, du kannst ja auch sagen, nein, für dich ist das kein Security, aber Operational Improvement. Das ist ja voll OK.

  30:35
Mhm, ja.

   30:51
die nächste Frage ist auch so, das geht in die gleiche Ding, ist auch eine Frage, was du da Security Ebene siehst. Bei Planetex, was du siehst, oder der hat jeder, was er vom Donkey Zugriff hat, im Prinzip ist ja genau das gleiche wie mit dem mit dem Security im im Ding da, da hast du im Prinzip das dasselbe.

  30:55
Mhm.
Genau, genau.

)   31:07
du kannst, du kannst alles.

   31:08
Aber wir haben aus Mangel an Optionen haben wir keine andere Chance.

)   31:12
Ja, die Sache ist rein kritisch könntest du halt auch in der Linux Welt einfach mehr pro Applikation einen User haben und diesen das dann Gruppen passiert also per se ist es halt in Security.

  31:21
Es liegt dann trotzdem noch auf Blindtext da.

)   31:24
Das ja, ja, okay, gut. Ja, wie sieht es ist ist ein acceptable? Das ist ein N und nicht A. Das muss ich auch ausbessern. Security Trade of.

  31:33
Mhm.
Mhm.
Thin surrence you suppricial complex.

   31:44
Ist ja die Frage, was irgendwas, wo man sich den Vergleich zum extern sieger Operator.

)   31:50
Jo.

   31:50
Oder ob man's aus Grund, ding, det weißt, was ich mein.

)   31:55
Ja, so wie so wie du die Frage auffällst.

   31:56
Mhm.
The agent has the potential to reduce operational complexity and secret management. Das muss man keinem komplett nahestellen, weil die Komplexität verringert sich ja nicht.

)   32:30
M.

  32:33
So gesehen, ich hab es auf auf Jo gestellt, weil es das Ausrollen vor ein.
Ach, aber ja, der Zugsegen.

   32:41
Der Einfache ist in Gitle gespeichert haben also.

  32:46
So gesehen, ja.

)   32:48
Der Initial Bootstrap ist definitiv komplex, aber dafür, wenn du es einmal konfiguriert hast, ist es einfacher. Das Gleiche kannst du ja sagen, auch es wäre so viel einfacher, ändert sich einfach jeden Ton, geht händisch zu warten, wenn es

   32:52
Yeah.

  32:52
Mhm.

)   33:00
Wenn es, wenn die Zeit da ist.

  33:02
Mhm, das Initial Setup ist leichter, wenn du es manuell einfach mal ausrollst, aber das.

   33:04
Yep.

)   33:04
M.
Ja, einmal machst du, aber wenn du so wie.

   33:08
Mhm.

  33:09
Aber das Operational Dings im Sinne von auf die Dauer ist es gescheiter anders ja.
Willst du es erfahren?
Witzig war der, weil weil ich gerade bei der Frage bin, the user facing interaction between the ESO and the Linux secret agent feels kohärent. Ja, vor vorletzte Frage, aber doch

)   35:30
Huh.

   35:31
Ist schon weit.

  35:38
es wäre eigentlich total witzig, wenn der, wenn der Linux Agent einfach die exakte Config Files wieder externen Secrets Operator nimmt.

)   35:47
Ja, das hab ich 'ne Zeit lang auch überlegt.
Ob er die genau gesagt gleich nimmt, das Sache ist, du musst halt etwas in und sowas anders anders unterbringen. Und das Problem ist, der Secret Operator kann erlaubt halt für für auch für für mehr als nur Templating.

  35:57
Ach so, ja, muss sowieso was anderes machen. Mhm.
Mhm, OK.

)   36:08
Demnach für den Prototypen geht das nicht, aber per se ist das etwas, was man noch aufnehmen kann in.

  36:13
Mhm.

)   36:16
Die Insets.

   36:22
Ist das ist relativ, also manche Fragen hast gestellt, die die was die eigentlich.
Musst du in in, in, in, wo du überlegen musst, in was für Richtung gehst.

)   36:34
OK, ja, wie gesagt, wenn es eine Frage gibt.

  36:35
Yeah.

   36:36
The workflow reduces effort in managing secrets.

)   36:41
Ja, das ist eine Frage. Du hast ihn kennengelernt. Siehst du per se jetzt einen Reduce in Effort? Ja oder nein?
Wenn du meinst, nein, du siehst per se keinen.

   36:52
Das Problem im Vergleich zu.

)   36:53
Canon Air?
Na, im Vergleich zu deiner aktuellen Lage, also wenn du keine eigene hast, dann halt manual, aber sonst halt zu dem, wie du es jetzt machen würdest.

   37:11
Weil die die die neue SSDLC Lösung halt egal was für Lösung das ist halt immer mehr Aufwand.

)   37:17
Yeah.
Per se, meine Arbeit geht auch davon aus, meine Arbeit geht. Ja, ich weiß, meine Arbeit geht auch per se davon aus, dass du einen Secret Store verwendest.

   37:20
Das hat aber nix, das hat aber nix mit deinem Agents zu tun.
Okay.

)   37:29
Vielleicht ist das noch wichtig zu erwähnen. Danke, nehme ich mir als Insert mit für die nächste Evaluierung.

   37:38
Weil weißt du, gegenüber KitLab ist natürlich jetzt ein komplettes ****** Situation.
Aber du musst doch sowas machen, geht nicht mehr.

  37:48
Yeah.

)   37:49
Ja.

   37:49
Als GitLab war natürlich einfacher.

)   37:52
Ja, aber halt noch mehr insecure.

   37:56
Ja klar, aber weniger Aufwand. Ist die Frage im Vergleich zu was?

)   37:57
Yeah.

  38:01
Das stimmt.

)   38:03
Hm.

   38:08
Aber jetzt vergleiche, indem dass man sagt, man da jetzt so wirklich die den Secret Volt oder was in die Applikation einbauen, dann ist sicherlich weniger offen. Da ist wieder ganz was anderes.

)   38:18
Ja, wie gesagt, das das größte Problem, warum ich sozusagen gegen das gehe, weil für mich ist das halt Blue Code, der nicht überwachbar ist, weil du halt nie weißt, ob das Rollieren in deiner Applikation dann wirklich funktioniert hat, weil sonst musst du das dann auch noch einbauen und ja.

   38:31
Da gibt es für mich ganz einen noch einen entscheidenden Nachteil, wenn du es in eine Applikation einbaust. Du tust das Segret Store aus, besser durch einen anderen. Dann musst du alle Applikationen umbauen, so du den Segret Store austauschen.

  38:40
Mhm.

)   38:43
Ja genau, du musst halt alle alle Credentials und alle Applikationen musst dann halt umbauen.

   38:44
Ist ein großes Ding.
Also, die rotation schedules are alignable with organizational constraints, tue ich bewusst auf normal, nur dass du es weißt, weil ich einfach noch keine Basis habe, wo ich das vergleichen kann.

)   39:03
Huh.

   39:04
Or it was, was, was, was more wirklich win.

  39:33
Was?

   39:35
Das Gleiche für die Workflow-Seams easy to integrate into the existing CICD-Pipelines. Das wissen wir nicht, wie wir es integrieren wollen.

)   39:42
Ja, aber darum seems.
Wäre es siems, wäre es einfach, diesen Kerl Befehl in.

  39:45
Teams.

)   39:51
Etwas einzubauen.
Weil das ist ja alles, was du machen musst. Das ist ein Kerbal Film mit entweder Post Delete.

   39:58
Das weiß ich nicht, deswegen sag ich dir da, da gehen wir in die Mitten, dann ist neutral. Ich kann dir das noch nicht sagen. Also, das haben wir letzte Woche diskutiert, wo wir mit Hafrox geredet haben, ob wir da hinkommen. Beziehungsweise, ob der User die oder über die URL veröffentlicht, über die URL veröffentlichen will ich das nicht. Da gehört jetzt auch dies URL, damit es auf Localhost bleibt. Da muss man erst mit *** reden, deswegen.

)   39:58
Auf.
OK.
Yeah.

   40:17
Da bin ich überhaupt noch nicht da.
Sicher, wie man das macht.

)   40:22
Yeah.
Aber per se, du musst ja nicht nur aus der aktuellen Position denken, sondern einfach, wenn würdest hättest du aufgrund der Security Architektur Architektur das Vertrauen, also wenn du gesagt hast, du hast das Vertrauen, dass der Internet facing ist.

   40:32
Mhm.

)   40:42
Aufgrund der Security Architektur, dann kannst, dann solltest du mental davon ausgehen, dass er ja im Internet hängt.

   40:49
Mhm.
Ich möcht, ich möcht das nur so, so quasi sagen, wenn da jetzt irgendwas ein wenig an die linke Seite runtergeht, das ist dann natürlich das negative Bewertung sehen für den Agent. Das möcht ich nur sagen.

)   40:53
Yeah.
Du bist du. Per se ist es keine Bewertung für mich. Es ist ein Experteninterview, wo ich dann rausschreiben werd, wie labern das ist, was ich gemacht habe. Dann werde ich in Future Work sagen.

   41:06
Na, passt alles gut.
Ne, weil die Antworten sollten dei dei dei Arbeit nicht schmälern, weil das ist oft teilweise ja die, die die jetzige Situation, die was teilweise Antworten noch nicht.
100% klar definieren lassen, sage ich es mal so.

)   41:23
Er versteht es.
Ja, du, alles, was wir reden, kann ich ja dann aufnehmen und darlegen, warum das ist und dann kann man sozusagen sagen.

   41:29
Mhm.

)   41:34
In den Limitations und in Fìr Gawinnen, was alles zu tun wäre, um es noch besser zu machen.

  41:38
Und da muss noch was.

   41:42
Weil die die Komplexität unserer Infrastruktur da kann ich da nichts dafür.

  41:47
Ja, ist ein normaler Punkt.

   41:50
Das musst du auf alle Fälle mitnehmen, dass manche Sachen einfach schwer zu betrachten beurteilt sind oder vielleicht ein bisschen anders ausfallen, als du vielleicht vorstellst. Aber das ist ja bedingt aufgrund von der jetzigen Infrastruktur, vielleicht.

)   41:51
Hm.
Ja, passt. Werde ich, werde ich mitnehmen.

   42:03
Und das bei uns, du wirst es zu Mark immer so, bei den anderen Firmen geht das halt viel leichter, aber ist mit euch eine komische Restriktion, ist eine Katastrophe.

)   42:12
Mhm.

  42:27
Ich habs auch so abgeschickt.

)   42:29
Dankeschön.

  42:31
Okay, auch.

   42:31
Ich schick sie auch mal ab.

)   42:33
So, jetzt haben wir dann noch ein paar Minütchen zum Reden über die offenen Fragen, die ich euch noch stellen möchte.

   42:44
Jetzt werden wir vorsichtig.

  42:45
Ja, ich glaube.

)   42:47
With it.
Und ihr überschätzt meine Fähigkeiten, mal wieder mal siehst.

   42:53
Weißt du, du musst uns jetzt dann noch einen Vertrag unterschreiben, dass du für das Phishing mit V jetzt keine bösen Absichten mit uns hast.

  43:00
Stimmt.

)   43:02
So.
Lest euch die Fragen durch und wenn es irgendeine gibt, ich kann sie hier vorlesen, wie ihr wollt.
Und wenn ihr irgendwo was dazu sagen wollt, dann freu ich mich und sonst?

  43:16
What do you see is the main advantment of the Leo secret agent?
Zum ersten Punkt, dass Sie.
Den gleichen Mechanismus wie für Open, also wie im OpenShift, auch für Linux nutzen kann.

   43:31
With the actually good.

  43:36
Wo ich momentan nicht dazu in der Lage bin.
Some support.

)   43:41
Versteh, du würdest.

   43:42
Ja, this is.

)   43:44
Mhm, das Coole ist, ich hab per se ja auch dann mit den 2 Open Shift Jungs noch ein Interview.

   43:46
Mhm.

)   43:51
View, die halt schon mehrmals als SIM-Cooperator gearbeitet haben und den gibt es halt coole Demos, das ist halt per se, die ist es halt nur einer der Sachen, genau ich.

  43:52
Mhm.

   43:54
Hm.
Ich weiß nicht, ob die schon mitgearbeitet haben, sie kennen den zumindestens.

  44:00
Mhm.

)   44:03
Ja, ja, also ich denke schon, wenn Sie ihn ausgerollt haben, werden Sie ihn verstanden haben, hoffe ich.

  44:10
Und und eben das.
Dass ich nicht gezwungen bin, nur um diese Features in OpenShift nutzen zu können, nach obenShift zu migrieren, sondern dass er doch entspannt die Migrationsphase quasi.

   44:21
Richtig, ja, das wollt ich.

)   44:22
Das
Dass ich genau einfach in hybriden Environments einen funktionierenden Workflow habe.

  44:28
Mhm.
Genau, kann schon.

   44:30
Und dass man relativ schnell die die Secrets aus einem Git rausbringen und äh.
Das in einem sicheren Store speichern können.

  44:37
In einem Secret Store, genau.

   44:40
Ohne die Applikationen zu ändern, und das ist der große Vorteil.

)   44:40
Gut.

  44:42
Genau, ohne die R stimmt. Das ist ein guter Punkt, ohne an die Applikation was schrauben zu müssen.

)   44:45
Ja, und rein kritisch könnte man halt auch Brownfield Applikationen, die man zugekauft hat, solange sie nur ein File haben, damit modernisieren und secret compliant machen. Genau. Gut, das haben wir eh schon geredet. Die zweite Frage.

  44:51
Mhm.

   44:53
Umm.
Genau.

  44:58
Mhm.

)   44:59
Das mit den haben wir auch schon beantwortet.

  45:04
****.

)   45:04
Und jemanden, in denen ihr den nicht verwenden würdet.

  45:10
In nicht Linux Umgebungen ach ha ha ha.

)   45:12
Mhm, ja, es kann ja sein, dass ihr sagt, dass ihr sagt, für so und so eine Applikation könnte ich es mir nicht vorstellen, diesen Agent zu verwenden. Wenn nicht, dann ist das ja auch OK, Das sind sozusagen Open Questions, sondern einfach nur ein bisschen.

  45:24
Mhm.

)   45:28
Das ist jetzt ein Weg.

   45:28
Hast du vorher schon dir selber die Antwort gegeben, dass man auch, wenn man eine Third Party Applikation hat, so lang so file hat, kannst du das abbilden und fertig.

)   45:34
Genau, du könntest rein theoretisch, ja, du könntest.
Genau, du könntest eindeutig auch, wie auf diesem Solution Build, könntest du auch deine Datenbanken sozusagen damit machen und alles andere so langsam file konsumiert. Genau, und da kommen wir auch schon. Wo könntet ihr sozusagen so einen Git-ups-Approach sehen, außer im

   45:39
Somit.
Mhm.

)   45:57
Im Secret Management.

  46:00
Ja, theoretisch könnte das die ganze Config so abbilden. Ja, weil, weil im Grunde, ob du das über also im Grunde ist ja die, die ich kann, ich könnte ja Config Map genauso pushen wie a wie a Secret im Grunde und hab dann

)   46:02
Ja.
Könntest.
Genau.
Genau.

  46:17
Im im Linux Bereich praktisch das das Look and Feel vom vom Openshift mehr oder weniger.

)   46:24
Ja, also man könnte per se überlegen, ob vielleicht Secret Management nicht das Ende sein muss, sondern ein Feature von Configuration Management damit sein könnte.

   46:25
Mhm.

  46:32
Mhm.
Mhm.
What would the agent need to demonstrate to be considered production rating? Ich glaube, da müsste man erst noch mal eben so was wie XML Escaping noch mal genauer anschauen. Jetzt müssen wir ob ob an OpenShift Sequence Operator A.

   46:50
Mhm.
They do you.

  46:54
dann müssen wir, nachdem wir die URL dann nach außen geben, das mit der Security noch mal beleuchten, rein vom von dem Bild her, was du gezeigt hast, wird es schon passen oder klingt es einmal richtig, aber das muss man sich noch mal extra beleuchten.

)   47:04
Yep.

  47:13
Umm.

   47:13
Ja, was, was da, was bei was auf alle Fälle offen ist, ist, dass den den Anstoßt, weil wie gesagt, die den Endpoint von außen, dass der Cenkin seinen Aufruf hat, das will ich ehrlich gesagt eher nicht.

  47:19
Mhm.
Genau.
Mhm.

   47:29
Das muss irgendwie über die den den Mechanismus des Deploys passieren. Beim Ding passiert es mit dem beim OpenShift bis jetzt mit dem externen Secret, dass es quasi getriggert wird, dass das mit Helm aufgespült wird.

)   47:38
Ja, aber per se per se hat der Open Shift halt einen external Endpoint.

  47:45
Ja, der Unterschied ist aber, dass dass du in OpenShift die klare Trennung hast zwischen dem dem Management Interface oder die Control-Pain von OpenShift und das ist ein eigener komplett eigener Endpoint.

   47:56
Yep.
The Secrets Connect.

  48:02
Und und dem Applikations URL erzeugsel und da muss muss man sich erst mal anschauen, wie das im im Linux Bereich funktionieren würde, weil das ja auf der gleichen Ebene quasi ist. Also infrastrukturell unterscheidest du nicht zwischen zwischen so eine ist Applikation.

)   48:07
Yeah, well.

   48:07
It pretty, I pretty, yeah.

)   48:13
Ja.

   48:14
Mhm, genau, I see.

  48:21
Uns URL und das ist diese Management URL vom vom Buildings Agent.

   48:23
Ja, ja, da denk beim beim OpenShift hast du den den Helm, deploy quasi und der deployed das external segrets und mit dem magst du den Connect Richtung des external segrets triggert dann quasi den Connect und ist da die Champions, den den Connect triggern Richtung Secret.
sondern das soll der, der, der deploy Vorgang. In dem Fall müsste es eigentlich, wenn man es richtig nimmt, das *** steuern. Dann passt das Ganze. Dann ist da was mit *** erst reden, wo ich jetzt eher von dem von dem Agent, den ich heute werde, den auch ausprobieren will oder oder den anschauen will oder

  48:46
Mhm.

   48:55
ob er das komplett aus dem Agent eigentlich wegkriegen will und komplett integrieren will in das ***. Das ist die andere Frage, weil er ja das vom Konf-Verzeichnis in das Konf-Kette den Lokal aus dem Deploy kopiert und anfänglich war er das so konzipiert, dass er das beim Deployvorgang
Wann hast du umkopiert, das dann quasi ersetzt?
Oder äh, die ganzen Passwörter damit ausliest und und und das Template rendert quasi.

  49:25
No.

   49:27
das muss jetzt mit dem *** reden, aber eigentlich gehört es entweder wirklich, dass man es aus, also als Age weg tut und rein tut in das, in den, dass das das Diplom macht. Oder es gehört rein, dass das ***

  49:27
Also, das das muss man sich.

   49:41
Quasi bei einem Start oder bei irgendwas oder das, äh, oder bei beim Deploy eigentlich das Ding dann aufruft vorher.
Aber da ist die Frage, was, was der *** will und des da haben wir, da haben wir jetzt weniger Einfluss.

)   49:53
Ja, aber dann.
Dann gehst du halt von, dann gehst du halt von dem ganzen Agent Konzept weg. Aber wie gesagt, was also sozusagen dann denk.

   50:00
richtig, aber das liegt jetzt halt zum Teil in unserer Hand, weil wenn es der *** jetzt mit der U. R. L. nicht integrieren will für uns, dann werden es wir nicht in der Art und Weise verwenden können. Das, weil also wir lösen uns

)   50:04
Ja.
Ja.
Ja, aber jetzt einmal.

   50:28
ausprobieren und mit *** auch noch ein C. O. Zoan und den das wirklich live am am am Ding Zoan, damit du es einfach mal siehst. Wir haben sie am *** sie installiert. Das heißt, da können wir sie ganz normal noch mal genau versorgen.

)   50:34
Yeah.
No.
Genau.

   50:44
Da mìsstest du die Internetseite mal ein bissel äh genauer aufschreiben, was man da aufrufen kann.

)   50:49
Passt, mach ich aber gut. Also, insofern wir nicht in uns, also das werde ich auch nicht mehr erwähnen, nicht in unserer Struktur. Kurz zusammengefasst, meint ihr, für Production Readiness sollte man noch die H-Cases genauer beleuchten, aber per se der Prototyp ist verwendbar.

   50:51
Die, die war nur cool, um in einer genauen Doku hätten Mhm.

)   51:06
Und löst das Problem vom dynamischen Secret Management auf Linux.

   51:09
Bis bis wann, bis wann wir uns die Frage genau beantwortet haben?

  51:14
Bis jetzt.

)   51:14
Jetzt.

   51:16
verdammt, weil, weil wir letzte Woche beschlossen haben, dass wir den Beruf of Concipits wirklich mal durchziehen. Das heißt, das war mal nett, wenn man das wirklich mal mit einem mit einem Donkey dann letzte Woche kurz mal ausprobiert, aber wo man das dann, besonders du, die Arbeit abgeben, sagen wir es mal so.

)   51:24
Yep.
Das sind.
Gib mir.
Per se, das ist nicht wichtig für meine Arbeit. Das ist einfach jetzt nur ein Bauchgefühl-Dings.

   51:34
OK, na passt. OK, ja, passt. Ja, ja, hundertprozentig. Also, ich könnte mir schon vorstellen.

)   51:37
Es ist, weil.
Weil per se per se, dass er das macht, was ich sage, das habe ich evaluiert aufgrund von von Tests und so weiter, das braucht ihr mir nicht sagen. Da geht es sozusagen einfach. Ihr seht das Konzept gesehen und den Prototypen und was ich gemacht habe und ja, genau.

   51:42
Mhm.
Okay, ja, passt.
Mhm.
Mhm.

  51:58
Das ist meine Freundin, die.
Genau das, was in der Frage steht, dass das ein einfacher Workflow ist, der im Grunde für beide von der vom Konzept her für beide gleich funktioniert.

   52:09
Yeah.

)   52:12
Put.
Ja, Limitations und Weaknesses haben wir eh schon besprochen, aber wenn ihr noch welche hinzufügen wollt, könnt ihr das gerne machen.

   52:17
Genau.
Ich glaube, das kann man im jetzigen Zeitpunkt nur schwer sagen, weil man es noch jetzt.

)   52:26
Ja.

   52:28
So genau so ausführlich verwendet haben, also vom Look and Feel müsste es passen, ja.

  52:28
Yeah.

)   52:31
Ja.

  52:32
In der Theorie her klings klings ganz okay, das ist die Separation of Secret.
Weg ist ja.

)   52:40
Also die die Fragen sind jetzt großartig, einfach nur da ich hab die schon beantwortet in der Like at Evaluation, um sozusagen da jetzt Punkte, die euch aufgefallen sind, strukturiere strukturiert am Ende noch mal abbilden zu können, dass ich mit denen arbeiten kann.

   52:45
Mhm.

  52:50
Mhm.

   52:51
Mhm.

  52:55
Also, Secret Authoring und Secret Delivery ist sicher gut und strengt, weil.

)   53:00
Das Schöne ist, Developer sehen ein Secret Value niemals mehr. Sie müssen nur noch den Secret Key wissen und können sich sozusagen selber ihre eigenen Configurations maintainen, weil sie überhaupt nichts von dem Store wissen müssen. Das Einzige, was sie wissen müssen, ist, wie das im Store heißt.

  53:05
Genau.
Genau.
Genau.

)   53:17
Und das kann man Ihnen ja freigeben oder das ist halt einmal definiert und aber genau, also ja, passt.

  53:17
Genau.

   53:17
Mhm.

)   53:25
Billy man.

   53:25
Wobei das das es gar nicht mehr sehen ist relativ, weil du du spielst das ja quasi eine Initial.

  53:25
Well, das der Deckel heute.

)   53:31
Ja, du spielst es am Server, aber per se muss sie muss sie keinen Zugriff zum Server haben.

   53:38
Der Developer wahrscheinlich nicht, ne.

)   53:40
Ja, genau darum geht es.

   53:42
Haben Sie aber momentan dieses noch irgendwie die Schwierigkeit, muss man sicherlich noch mal nacharbeiten.

)   53:45
Ja, das.

  53:48
Ja.
Aber wenn Sie Zugriff auf den Server haben, dann hätten Sie egal, also egal was du für einen Approach hast.

   53:50
Yep.

)   53:52
Yeah, aber aber.

   53:56
Mhm.

)   53:56
Nicht unbedingt, wenn Sie jetzt sozusagen am Server sind und sozusagen nur auf Lockdirectories mit Ihrem Elder User berechtigt wären, dann würde es Ihnen nichts bringen.

  54:05
Ja, ja, da am Server, das stimmt schon, aber ich meine mit mit mit dem Applikationsuser am Server, das ist per se eigentlich eh ein ein ein No-Go.

   54:07
Mhm.

)   54:07
Bell.
Yeah.
And and yeah.

   54:13
Genau, aber du musst im Prinzip, äh, wärst das Skript, wenn man machen Mìsst, dass halt irgendwer irgendwo als, äh, secret aufspielen kann. Das heißt, das muss man so irgendwie machen.

  54:18
Yeah.
Mhm.

   54:23
Das ist ja die Frage, wer das Skript dann ausführt, weil wenn jetzt die Entwickler das wirklich machen sollten, dass sie eine neue Umgebung aufsetzen, dann müssen sie halt die neue Secrets irgendwie aufspielen können. Weil sie sehen ja die Secret dann vielleicht auch nicht, weil das Skript vielleicht das Mama kann das Skript wieder ändern. Ja, O. K.
Man kann es natürlich sehen, ob er.

)   54:41
Ja, aber wie?

  54:41
Ja.

   54:41
Es wird nicht gestored. Das ist jetzt zumindest das.

)   54:44
Ja, oder oder man gibt Entwicklern halt einfach generell nur die Möglichkeit, den Store zu sehen und alle Secret Values, die sie haben müssen, muss man einen DevOps per Ticket machen und der setzt dir die Werte auf und dann kannst du damit arbeiten. Aber das ist ja, das ist ja Organisational Boundaries, das hat ja nichts mit meinem Agent.

   54:54
Genau das wissen wir ja momentan noch nicht.

  54:58
Mhm.

)   54:59
Und dem Workflow zu tun.

  55:01
Git-based approach, align with your existing Git-Obsource E. I. C. D. Practices, ja, ich glaube, da haben wir, haben wir auf unserer Seite noch Aufholbedarf, aber an und für sich wird man uns schon schon einbinden können. What Improvements would be required

   55:11
Mhm, It's Claudia.

)   55:17
Die Frage ist, die haben wir auch schon beantwortet.

  55:18
Before you documented production.

)   55:22
Yeah.

  55:23
Yeah, what would prevent the organization from adopting this workflow? Ja, wir müssen es ausprobieren. OK, genau, ja.

)   55:28
Organizational constraints, yeah.

   55:34
Ich will, wie gesagt, das muss man anschauen oder wirklich, dass man andere Leute auch beobachten muss oder befragen muss und das freigeben muss, damit.

)   55:37
Wunderbar.
Gut.

) Transkription beendet
